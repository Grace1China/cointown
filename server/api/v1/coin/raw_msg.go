package coin

import (
	"fmt"
	"time"

	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"

	coinReq "github.com/Grace1China/cointown/server/model/coin/request"
	"github.com/Grace1China/cointown/server/model/common/response"
	"github.com/Grace1China/cointown/server/task"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RawMsgApi struct{}

// CreateRawMsg 创建原始消息
// @Tags RawMsg
// @Summary 创建原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.RawMsg true "创建原始消息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /rawmsg/createRawMsg [post]
func (rawmsgApi *RawMsgApi) CreateRawMsg(c *gin.Context) {
	var rawmsg coin.RawMsg
	err := c.ShouldBindJSON(&rawmsg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//1 存入消息 y
	//2 存入频道，或者更新频道
	//3 如果是更新频道，取出template 并解析出来代币信息
	//  3.1 存入代币
	//4 如果是新建频道，需要提示去建立模版
	//5 在仪表盘中显示需要建立模板的频道和需要跟进的解析错误的消息

	err = rawmsgService.CreateRawMsg(&rawmsg)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	rawmsg, err = rawmsgService.GetRawMsg(fmt.Sprintf("%d", rawmsg.ID))
	right, coin := tempParseService.ParseCoinsAuto(rawmsg)
	fmt.Println(right, coin)
	if right {
		//如果此币地址已经出现过，只对币信息做更新
		coin1, err := coinsService.GetCoinsByAddr(coin.ContractAddress)
		if err != nil {
			now := time.Now()
			coin.PriceTestTime1 = now.Add(5 * time.Minute)
			coin.PriceTestTime2 = now.Add(15 * time.Minute)
			coin.PriceTestTime3 = now.Add(30 * time.Minute)
			coin.ChatId = rawmsg.ChatId
			coin.MsgId = rawmsg.ID
			if coin.Price <= 0 {
				url1 := fmt.Sprintf("https://api.solanagateway.com/api/v1/pumpfun/price?mint=%s", coin.ContractAddress)
				url2 := fmt.Sprintf("https://api.solanagateway.com/api/v1/raydium/price?mint=%s", coin.ContractAddress)
				chanPrice1 := task.FetchPriceAsync(url1)
				chanPrice2 := task.FetchPriceAsync(url2)
				pair := <-chanPrice1
				pair2 := <-chanPrice2
				// price111, err := task.FetchPrice(url1)

				if pair.Err == nil {
					coin.Price = pair.Price
				} else if pair2.Err == nil {
					coin.Price = pair2.Price
				}
				err = coinsService.CreateCoins(&coin)
				if err != nil {
					global.GVA_LOG.Error("创建Coin失败!", zap.Error(err))
					response.FailWithMessage("创建Coin失败:"+err.Error(), c)
					return
				}
			}
		}
		coin.ID = coin1.ID
		coin.MsgId = rawmsg.ID
		// coin.ChatId = rawmsg.ChatId
		err = coinsService.UpdateCoins(coin)
		if err != nil {
			global.GVA_LOG.Error("更新Coin失败!", zap.Error(err))
			response.FailWithMessage("更新Coin失败:"+err.Error(), c)
			return
		}
		response.OkWithMessage("创建成功", c)
	} else {
		global.GVA_LOG.Info(fmt.Sprintf("发现了新消息 id=%d ChatId = %s TopicId =%s ", rawmsg.ID, rawmsg.ChatId, rawmsg.TopicId))
		global.GVA_LOG.Info(rawmsg.MessageText)
		rawmsg.IsNew = true
		rawmsgService.UpdateRawMsg(rawmsg)
		response.OkWithMessage("发现了新模版", c)
	}

}

// DeleteRawMsg 删除原始消息
// @Tags RawMsg
// @Summary 删除原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.RawMsg true "删除原始消息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /rawmsg/deleteRawMsg [delete]
func (rawmsgApi *RawMsgApi) DeleteRawMsg(c *gin.Context) {
	ID := c.Query("ID")
	err := rawmsgService.DeleteRawMsg(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRawMsg7D DeleteRawMsg7D
// @Tags RawMsg
// @Summary DeleteRawMsg7D
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.RawMsg true "删除原始消息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /rawmsg/DeleteRawMsg7D [post]
func (rawmsgApi *RawMsgApi) DeleteRawMsg7D(c *gin.Context) {
	err := rawmsgService.DeleteRawMsg7D()
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRawMsgByIds 批量删除原始消息
// @Tags RawMsg
// @Summary 批量删除原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /rawmsg/deleteRawMsgByIds [delete]
func (rawmsgApi *RawMsgApi) DeleteRawMsgByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := rawmsgService.DeleteRawMsgByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRawMsg 更新原始消息
// @Tags RawMsg
// @Summary 更新原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.RawMsg true "更新原始消息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /rawmsg/updateRawMsg [put]
func (rawmsgApi *RawMsgApi) UpdateRawMsg(c *gin.Context) {
	var rawmsg coin.RawMsg
	err := c.ShouldBindJSON(&rawmsg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rawmsgService.UpdateRawMsg(rawmsg)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// UpdateRawMsg UpdateRawMsgStatusAll
// @Tags RawMsg
// @Summary UpdateRawMsgStatusAll
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.RawMsg true "更新原始消息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /rawmsg/UpdateRawMsgStatusAll [post]
func (rawmsgApi *RawMsgApi) UpdateRawMsgStatusAll(c *gin.Context) {
	// var rawmsg coin.RawMsg
	// err := c.ShouldBindJSON(&rawmsg)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	err := rawmsgService.UpdateRawMsgStatusAll()
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRawMsg 用id查询原始消息
// @Tags RawMsg
// @Summary 用id查询原始消息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query coin.RawMsg true "用id查询原始消息"
// @Success 200 {object} response.Response{data=coin.RawMsg,msg=string} "查询成功"
// @Router /rawmsg/findRawMsg [get]
func (rawmsgApi *RawMsgApi) FindRawMsg(c *gin.Context) {
	ID := c.Query("ID")
	rerawmsg, err := rawmsgService.GetRawMsg(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rerawmsg, c)
}

// GetRawMsgList 分页获取原始消息列表
// @Tags RawMsg
// @Summary 分页获取原始消息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query coinReq.RawMsgSearch true "分页获取原始消息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /rawmsg/getRawMsgList [get]
func (rawmsgApi *RawMsgApi) GetRawMsgList(c *gin.Context) {
	var pageInfo coinReq.RawMsgSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ChatId := c.Query("ChatId")
	id := c.Query("id")
	pageInfo.ChatId = ChatId
	pageInfo.Id = id

	list, total, err := rawmsgService.GetRawMsgInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetRawMsgPublic 不需要鉴权的原始消息接口
// @Tags RawMsg
// @Summary 不需要鉴权的原始消息接口
// @accept application/json
// @Produce application/json
// @Param data query coinReq.RawMsgSearch true "分页获取原始消息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rawmsg/getRawMsgPublic [get]
func (rawmsgApi *RawMsgApi) GetRawMsgPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的原始消息接口信息",
	}, "获取成功", c)
}
