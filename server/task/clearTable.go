package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/common"
	"github.com/Grace1China/cointown/server/service"

	"gorm.io/gorm"
)

// @author: [songzhibin97](https://github.com/songzhibin97)
// @function: ClearTable
// @description: 清理数据库表数据
// @param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
// @return: error
var (
	rawmsgService    = service.ServiceGroupApp.CoinServiceGroup.RawMsgService
	coinsService     = service.ServiceGroupApp.CoinServiceGroup.CoinsService
	channelService   = service.ServiceGroupApp.CoinServiceGroup.ChannelService
	tempParseService = service.ServiceGroupApp.CoinServiceGroup.TempParseService
)

func ClearTable(db *gorm.DB) error {
	var ClearTableDetail []common.ClearDB

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "sys_operation_records",
		CompareField: "created_at",
		Interval:     "2160h",
	})

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "jwt_blacklists",
		CompareField: "created_at",
		Interval:     "168h",
	})

	if db == nil {
		return errors.New("db Cannot be empty")
	}

	for _, detail := range ClearTableDetail {
		duration, err := time.ParseDuration(detail.Interval)
		if err != nil {
			return err
		}
		if duration < 0 {
			return errors.New("parse duration < 0")
		}
		err = db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", detail.TableName, detail.CompareField), time.Now().Add(-duration)).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func GetPrice() error {
	// coin.CoinsService
	t := time.Now()
	coins, err := coinsService.GetCoinsByTimeDue(t)
	if err != nil {
		return err
	}
	for i, coin := range coins {
		url1 := fmt.Sprintf("https://api.solanagateway.com/api/v1/pumpfun/price?mint=%s", coin.ContractAddress)
		url2 := fmt.Sprintf("https://api.solanagateway.com/api/v1/raydium/price?mint=%s", coin.ContractAddress)
		// fmt.Println(url1, url2)
		fmt.Println("GetPrice ", i)
		if coin.PriceTestTime1.Before(t) && coin.PriceTest1 == 0.0 {
			//取第一个价格时，如果币的原始价格是0 -1 -2 就也把它改成PriceTest1
			if price, err := FetchPrice(url1); err == nil {
				coin.PriceTest1 = price
				if coin.Price <= 0 {
					coin.Price = coin.PriceTest1
				}
				fmt.Println("GetPrice1", i, price)
			} else if price, err := FetchPrice(url2); err == nil {
				coin.PriceTest1 = price
				if coin.Price <= 0 {
					coin.Price = coin.PriceTest1
				}
				fmt.Println("GetPrice1", i, price)
			} else {
				if coin.PriceTest1 == 0 {
					coin.PriceTest1 = -1
				}
				if coin.PriceTest1 == -1 {
					coin.PriceTest1 = -2
				}
				if coin.PriceTest1 == -2 {
					coin.PriceTest1 = -100
				}
				fmt.Println("GetPrice1 ", i, "not found", url1, url2)
			}
			if coin.Price > 0 && coin.PriceTest1 > 0 {
				coin.PriceTestWave1 = (coin.PriceTest1 - coin.Price) / coin.Price
			}
		}
		if coin.PriceTestTime2.Before(t) && coin.PriceTest2 == 0.0 {
			if price, err := FetchPrice(url1); err == nil {
				coin.PriceTest2 = price
				fmt.Println("GetPrice2", i, price)
			} else if price, err := FetchPrice(url2); err == nil {
				coin.PriceTest2 = price
				fmt.Println("GetPrice2", i, price)
			} else {
				if coin.PriceTest2 == 0 {
					coin.PriceTest2 = -1
				}
				if coin.PriceTest2 == -1 {
					coin.PriceTest2 = -2
				}
				if coin.PriceTest2 == -2 {
					coin.PriceTest2 = -100
				}
				fmt.Println("GetPrice2 ", i, "not found", url1, url2)
			}
			if coin.Price > 0 && coin.PriceTest2 > 0 {
				coin.PriceTestWave2 = (coin.PriceTest2 - coin.Price) / coin.Price
			}
		}
		if coin.PriceTestTime3.Before(t) && coin.PriceTest3 == 0.0 {
			if price, err := FetchPrice(url1); err == nil {
				coin.PriceTest3 = price
				fmt.Println("GetPrice3", i, price)
			} else if price, err := FetchPrice(url2); err == nil {
				coin.PriceTest3 = price
				fmt.Println("GetPrice3", i, price)
			} else {
				if coin.PriceTest3 == 0 {
					coin.PriceTest3 = -1
				}
				if coin.PriceTest3 == -1 {
					coin.PriceTest3 = -2
				}
				if coin.PriceTest3 == -2 {
					coin.PriceTest3 = -100
				}
				fmt.Println("GetPrice3 ", i, "not found", url1, url2)
			}
			if coin.Price > 0 && coin.PriceTest3 > 0 {
				coin.PriceTestWave3 = (coin.PriceTest3 - coin.Price) / coin.Price
			}
		}
		coinsService.UpdateCoins(coin)
	}
	return fmt.Errorf("no price found")
}

// FetchPrice 发起 HTTP 请求并解析 JSON 响应以获取价格。
func FetchPrice(url string) (float64, error) {
	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("%v -- %s", err, url))
		return 0, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		global.GVA_LOG.Error(fmt.Sprintf("request failed with status %d: %s --%s", resp.StatusCode, body, url))
		return 0, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, body)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("error decoding JSON response: %v --%s", err, url))
		return 0, fmt.Errorf("error decoding JSON response: %v", err)
	}

	// 类型断言
	priceStr, ok := data["priceInUSD"].(string)
	if !ok {
		body, _ := io.ReadAll(resp.Body)
		global.GVA_LOG.Error(fmt.Sprintf("no priceInUSD found in API response: %s --%s", body, url))
		return 0, fmt.Errorf("no priceInUSD found in API response: %s", body)
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		global.GVA_LOG.Error(fmt.Sprintf("no priceInUSD found in API response: %s --%s", body, url))
		return 0, fmt.Errorf("no priceInUSD found in API response: %s", body)
	}

	return price, nil
}

// 定义一个 channel 用来接收网络请求的结果
func FetchPriceAsync(url string) chan struct {
	Price float64
	Err   error
} {
	resultCh := make(chan struct {
		Price float64
		Err   error
	})

	go func() {
		resp, err := http.Get(url)
		// if err != nil {
		// 	resultCh <- fmt.Sprintf("Error: %v", err)
		// 	return
		// }
		if err != nil {
			// global.GVA_LOG.Error(fmt.Sprintf("%v -- %s", err, url))
			resultCh <- struct {
				Price float64
				Err   error
			}{0, err}
		}
		defer resp.Body.Close()

		var data map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			resultCh <- struct {
				Price float64
				Err   error
			}{0, err}
		}

		// 类型断言
		priceStr, ok := data["priceInUSD"].(string)
		if !ok {
			// body, _ := io.ReadAll(resp.Body)
			resultCh <- struct {
				Price float64
				Err   error
			}{0, err}
		}
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			resultCh <- struct {
				Price float64
				Err   error
			}{0, err}
		}

		resultCh <- struct {
			Price float64
			Err   error
		}{price, err}
	}()
	// pair := <-resultCh
	// return pair.price, pair.err
	return resultCh
}
