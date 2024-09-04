package coin

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"reflect"

	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"
	"go.uber.org/zap"
	// "github.com/Grace1China/cointown/server/model/coin/ChannelService"
)

type TempParseService struct {
	channelService ChannelService
}

// 定义一个函数类型
type Func func(coin.RawMsg) (bool, coin.Coins)

func getValue(pat string, data string, field string) (ret string, err error) {
	reg := regexp.MustCompile(pat)
	mat := reg.FindStringSubmatch(data)
	global.GVA_LOG.Info(field+" matched", zap.Any("mat", mat))
	// global.GVA_LOG.Info(fmt.Sprintf("pat %s field %s", pat, field), zap.Any("strinig", data))
	if len(mat) > 1 {
		// fmt.Println(field+":", mat[1])
		// global.GVA_LOG.Info(field + ":" + mat[1])
		ret = mat[1]
		return ret, nil
	} else {
		// fmt.Println(field + " no info")
		// global.GVA_LOG.Info(field + " no info")
		return "", errors.New(field + " no info")
	}
}

// func getValueArr(pat string, data string, field string) (ret []string) {
// 	reg := regexp.MustCompile(pat)
// 	mat := reg.FindStringSubmatch(data)
// 	if len(mat) > 1 {
// 		fmt.Println(field+":", mat[1])
// 		global.GVA_LOG.Info(field + ":" + mat[1])
// 		ret = mat[1:len(mat)]
// 		return
// 	} else {
// 		fmt.Println(field + " no info")
// 		global.GVA_LOG.Info(field + " no info")
// 		return
// 	}
// }

// func getValue2(pat string, data string, field string) (string, string) {
// 	reg := regexp.MustCompile(pat)
// 	mat := reg.FindStringSubmatch(data)
// 	if len(mat) > 1 {
// 		fmt.Println(field+":", mat[1])
// 		global.GVA_LOG.Info(field + ":" + mat[1])
// 		// ret = mat[1]
// 		return mat[1], mat[2]
// 	} else {
// 		fmt.Println(field + " no info")
// 		global.GVA_LOG.Info(field + " no info")
// 		return "", ""
// 	}
// }

// func getValue3(pat string, data string, field string) (string, string, string) {
// 	reg := regexp.MustCompile(pat)
// 	mat := reg.FindStringSubmatch(data)
// 	if len(mat) > 1 {
// 		fmt.Println(field+":", mat[1])
// 		global.GVA_LOG.Info(field + ":" + mat[1])
// 		return mat[1], mat[2], mat[3]
// 	} else {
// 		fmt.Println(field + " no info")
// 		global.GVA_LOG.Info(field + " no info")
// 		return "", "", ""
// 	}
// }

// setFieldByString 通过字段名字符串设置结构体字段的值
func setFieldByString(obj interface{}, fieldName string, value interface{}) (err error) {
	// 获取对象的反射值
	val := reflect.ValueOf(obj)

	// 检查对象是否是指针
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// 获取字段
	field := val.FieldByName(fieldName)

	if !field.IsValid() {
		return fmt.Errorf("field '%s' does not exist", fieldName)
	}

	// 设置字段值
	if field.CanSet() {
		// 获取值的反射值
		val := reflect.ValueOf(value)

		// 检查值的类型是否匹配
		if val.Type().ConvertibleTo(field.Type()) {
			field.Set(val.Convert(field.Type()))
		} else if field.Type().Kind() == reflect.Uint {
			uintV, err := strconv.ParseUint(value.(string), 10, 64)
			if err == nil {
				field.SetUint(uintV)
			} else {
				return fmt.Errorf("value  %s does not conv to unit ", value)
			}
		} else {
			return fmt.Errorf("value type %T does not match field type %s ", val.Type(), field.Type())
		}
	} else {
		return fmt.Errorf("field %s can not set", fieldName)
	}
	return
}
func (tempParseService *TempParseService) ParseCoinsAuto(rawmsg coin.RawMsg) (right bool, coin coin.Coins) {
	channels, total, err := tempParseService.channelService.GetChannelByChlTopicID(rawmsg.ChatId, rawmsg.TopicId)
	j := 0
	if err == nil && total > 0 {
		// global.GVA_LOG.Info("ParseCoinsAuto channels:", zap.Int64("total", total))
		for index, value := range channels {
			j++
			// fmt.Println("Element at index", index, "is", value)
			// global.GVA_LOG.Info("ParseCoinsAuto channel:", zap.Int("index", index), zap.Any("channel", value))
			var jsonData map[string]interface{}
			err := json.Unmarshal([]byte(value.Template), &jsonData)
			if err != nil {
				global.GVA_LOG.Error("err", zap.Error(err))
				break
			}
			// global.GVA_LOG.Info("2 ParseCoinsAuto channel:", zap.Int("index", index), zap.Any("channel", value))
			allRight := true
			for key, value := range jsonData {
				valueStr, ok := value.(string)
				if ok {
					global.GVA_LOG.Info("准备解析", zap.String("key", key), zap.String("expr", string(valueStr)))
					retValue, err := getValue(valueStr, strings.ReplaceAll(rawmsg.MessageText, "`", "<addr>"), key)
					if err != nil {
						global.GVA_LOG.Info("解析发生错误", zap.Error(err))
						global.GVA_LOG.Error("解析发生错误", zap.Error(err))
						allRight = false
						break
					}
					err = setFieldByString(&coin, key, retValue)
					if err != nil {
						global.GVA_LOG.Info("设置发生错误", zap.Error(err))
						global.GVA_LOG.Error("设置发生错误", zap.Error(err))
						allRight = false
						break
					}
					global.GVA_LOG.Info("解析结果", zap.String("key", key), zap.String("result", retValue))
				}
			}
			global.GVA_LOG.Info("3 ParseCoinsAuto channel:", zap.Int("index", index), zap.Any("channel", value))

			//当这个channel的模板字段全部解析正确的时候就算是命中，但以后可能对于模版字段需要加一个必填标识.只需要命中所有必填项目后就算命中
			if allRight {
				global.GVA_LOG.Info("4 ParseCoinsAuto all field parsed.", zap.Any("coin", coin))
				value.Hits = value.Hits + 1
				now1 := time.Now()
				value.HitAt = &now1
				tempParseService.channelService.UpdateChannel(value)
				return true, coin
			}
		}
		// global.GVA_LOG.Info("5 ParseCoinsAuto channel:", zap.Int("j", j), zap.Any("total", total))
		// if j == int(total) {
		//如果找完所用的模板，那么说明没找到全部匹配好模板
		global.GVA_LOG.Info(fmt.Sprintf("解析完频道没有模板命中! ChatId:%s TopicId:%s j:%d total:%d", rawmsg.ChatId, rawmsg.TopicId, j, total))
		return false, coin
		// }
	} else {
		global.GVA_LOG.Error(fmt.Sprintf("获取频道失败! ChatId:%s TopicId:%s j:%d total:%d", rawmsg.ChatId, rawmsg.TopicId, j, total), zap.Error(err))
		return false, coin
	}

}

// func (tempParseService *TempParseService) ParseCoins(rawmsg coin.RawMsg) (right bool, coin coin.Coins) {
// 	// 创建一个 map，键为字符串，值为函数

// 	// 1 来了消息，用模版一个一个的解析，找到第一个全部中 或者中得最多 或者至少有两tokenName和addr

// 	functionMap := map[string]Func{
// 		"-1002050228063_137318": tempParseService.getCoin1,
// 		"-1002050228063_137320": tempParseService.getCoin1,
// 		"-1002050228063_137324": tempParseService.getCoin1,
// 		"-1002205321414_":       tempParseService.getCoin2,
// 	}
// 	var index = fmt.Sprintf(`%s_%s`, rawmsg.ChatId, rawmsg.TopicId)
// 	fmt.Println("parsed ====================>:" + index)
// 	global.GVA_LOG.Info("parsed ====================>:" + index)

// 	if functionMap[index] != nil {
// 		right, coin = functionMap[index](rawmsg)
// 		if right {
// 			fmt.Println(index, "parsed right:"+strconv.FormatBool(right))
// 			global.GVA_LOG.Info(index + " parsed right:" + strconv.FormatBool(right))
// 		}
// 	}
// 	return
// }

// func (tempParseService *TempParseService) getCoin1(rawmsg coin.RawMsg) (right bool, coin coin.Coins) {
// 	data := strings.ReplaceAll(rawmsg.MessageText, "`", "<addr>")
// 	fmt.Println("====================================>getCoin1", data)
// 	global.GVA_LOG.Info(fmt.Sprintf("====================================>getCoin1: %s", data))
// 	coin.TokenName = getValue(`\*\*Name\*\*:([A-Za-z\s&-]+)`, data, "coin.TokenName")
// 	coin.ContractAddress = getValue(`<addr>(.*?)<addr>`, data, "coin.ContractAddress")
// 	coin.MarketCap = getValue(`\*\*Market Cap\*\*:(\$\d+)`, data, "coin.MarketCap")
// 	//第二种模板
// 	if coin.TokenName == "" || coin.ContractAddress == "" || coin.MarketCap == "" {
// 		coin.TokenName, coin.ContractAddress, coin.MarketCap = "", "", ""

// 		fmt.Println("=======================> template 2")
// 		global.GVA_LOG.Info("=======================> template 2")

// 		coin.TokenName = getValue(`\*\*(.*?)\*\* (\(.*?\))`, data, "coin.TokenName")
// 		coin.ContractAddress = getValue(`<addr>(.*?)<addr>`, data, "coin.ContractAddress")
// 		coin.Progress = getValue(`Progress: (\d+%)`, data, "coin.Progress")
// 		coin.Balance = getValue(`Balance: (\d+) SOL`, data, "coin.Balance")
// 		coin.MarketCap = getValue(`MarketCap: (\$\d+)`, data, "MarketCap")
// 		coin.CommentsCount = getValue(`Comments: (\d+)`, data, "CommentsCount")
// 		coin.DevHolding = getValue(`Dev Holding: (\d+%)`, data, "DevHolding")
// 	}
// 	if coin.TokenName == "" || coin.ContractAddress == "" || coin.MarketCap == "" {
// 		// Alerts: GIL `8JJ6HeTRWWyCbmAzM6egj8K3xHMptAp8kdunA4khpump` (click to copy) Address Value 🟢 2dY4X4 0.5Sol 🟢 AaHBhD 1.23Sol Smart Buyers: 2
// 		global.GVA_LOG.Info("=======================> template 3")
// 		coin.TokenName, coin.ContractAddress, coin.MarketCap = "", "", ""
// 		coin.TokenName = getValue(`Alerts:\n\s([a-zA-Z\-\d+]+)`, data, "coin.TokenName")
// 		coin.ContractAddress = getValue(`<addr>(.*?)<addr>`, data, "coin.ContractAddress")
// 		smartBuyer := getValue(`Smart Buyers: (\d+)`, data, "coin.SmartBuyer")
// 		smartBuyer1, err := strconv.ParseUint(smartBuyer, 10, 64)
// 		if err != nil {
// 			global.GVA_LOG.Info(fmt.Sprintf("parse smartBuyer:%s", smartBuyer))
// 			global.GVA_LOG.Error("SmartBuyer转换失败", zap.Error(err))
// 		} else {
// 			coin.SmartBuyer = uint(smartBuyer1)
// 			global.GVA_LOG.Info(fmt.Sprintf("parse smartBuyer:%s  %d", smartBuyer, coin.SmartBuyer))
// 		}
// 	}
// 	if coin.TokenName == "" || coin.ContractAddress == "" || coin.SmartBuyer == 0 {
// 		global.GVA_LOG.Info("parse false")
// 		return false, coin
// 	}
// 	global.GVA_LOG.Info("parse right")

// 	return true, coin
// }

// func (tempParseService *TempParseService) getCoin2(rawmsg coin.RawMsg) (right bool, coin coin.Coins) {

// 	data := strings.ReplaceAll(rawmsg.MessageText, "`", "<addr>")
// 	fmt.Println("======================================>getCoin2", "data")

// 	coin.TokenName = getValue(`代币简称: ([a-z|A-Z|\d+-]*)`, data, "coin.TokenName")
// 	coin.ContractAddress = getValue(`<addr>(.*?)<addr>`, data, "coin.ContractAddress")
// 	coin.MarketCap = getValue(`总量: (\d+)`, data, "coin.MarketCap")
// 	coin.CreatedTime = getValue(`创建时间: (\d+-\d+-\d+ \d+:\d+:\d+)`, data, "coin.CreatedTime")
// 	coin.Holder = getValue(`持有人: \*\*(.*?)\*\*`, data, "coin.Holder")
// 	coin.Top10 = getValue(`Top10持仓: \*\*(.*?) \*\*`, data, "coin.Top10")
// 	coin.Progress = getValue(`进度: \*\*\[.*?\] (\d+%)\*\*`, data, "coin.Progress")
// 	Price111 := getValue(`币价: \*\*(.*?)\*\*`, data, "coin.Price")
// 	coin.Price, _ = strconv.ParseFloat(Price111, 64)
// 	// 交易量: 1mtx: 69 | 5mtx: 483 | 1htx: 485
// 	coin.TradeVol1m, coin.TradeVol5m, coin.TradeVol1h = getValue3(`交易量: 1mtx: (.*?) \| 5mtx: (.*?) \| 1htx: (\d+)`, data, "coin.TradeVol")
// 	coin.BuyVol1m, coin.BuyVol5m, coin.BuyVol1h = getValue3(`买入: 5mtx: (.*?) \| 1htx: (.*?) \| 6htx: (\d+)`, data, "coin.BuyVol1m")
// 	coin.SaleVol1m, coin.SaleVol5m, coin.SaleVol1h = getValue3(`卖出: 5mtx: (.*?) \| 1htx: (.*?) \| 6htx: (\d+)`, data, "coin.SaleVol1m")
// 	coin.Circulating, coin.Liq = getValue2(`流通市值: \*\*(.*?)\*\* | 流动性: \*\*(.*?)\*\*`, data, "coin.Circulating")
// 	coin.CreatorBal = getValue(`余额: \*\*(.*?)\*\*`, data, "coin.CreatorBal")
// 	coin.CreatorBalPer = getValue(`持有: \*\*(.*?)\*\*`, data, "coin.CreatorBalPer")
// 	coin.MintCoins = getValue(`创建代币: (\d+) 次`, data, "coin.MintCoins")
// 	coin.RecommendAddr = getValue(`推地址:\*\*(.*?)\*\* `, data, "coin.RecommendAddr")
// 	coin.BlackList = getValue(`黑名单: \*\*(.*?)\*\*`, data, "coin.BlackList")

// 	if coin.TokenName == "" || coin.ContractAddress == "" || coin.MarketCap == "" {
// 		return false, coin
// 	}

// 	return true, coin
// }
