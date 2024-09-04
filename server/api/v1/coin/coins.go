package coin

import (
	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"
	coinReq "github.com/Grace1China/cointown/server/model/coin/request"
	"github.com/Grace1China/cointown/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CoinsApi struct{}

// CreateCoins 创建币
// @Tags Coins
// @Summary 创建币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.Coins true "创建币"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /coins/createCoins [post]
func (coinsApi *CoinsApi) CreateCoins(c *gin.Context) {
	var coins coin.Coins
	err := c.ShouldBindJSON(&coins)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = coinsService.CreateCoins(&coins)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCoins 删除币
// @Tags Coins
// @Summary 删除币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.Coins true "删除币"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /coins/deleteCoins [delete]
func (coinsApi *CoinsApi) DeleteCoins(c *gin.Context) {
	ID := c.Query("ID")
	err := coinsService.DeleteCoins(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCoinsByIds 批量删除币
// @Tags Coins
// @Summary 批量删除币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /coins/deleteCoinsByIds [delete]
func (coinsApi *CoinsApi) DeleteCoinsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := coinsService.DeleteCoinsByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCoins 更新币
// @Tags Coins
// @Summary 更新币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.Coins true "更新币"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /coins/updateCoins [put]
func (coinsApi *CoinsApi) UpdateCoins(c *gin.Context) {
	var coins coin.Coins
	err := c.ShouldBindJSON(&coins)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = coinsService.UpdateCoins(coins)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCoins 用id查询币
// @Tags Coins
// @Summary 用id查询币
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query coin.Coins true "用id查询币"
// @Success 200 {object} response.Response{data=coin.Coins,msg=string} "查询成功"
// @Router /coins/findCoins [get]
func (coinsApi *CoinsApi) FindCoins(c *gin.Context) {
	ID := c.Query("ID")
	recoins, err := coinsService.GetCoins(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recoins, c)
}

// GetCoinsList 分页获取币列表
// @Tags Coins
// @Summary 分页获取币列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query coinReq.CoinsSearch true "分页获取币列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /coins/getCoinsList [get]
func (coinsApi *CoinsApi) GetCoinsList(c *gin.Context) {
	var pageInfo coinReq.CoinsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := coinsService.GetCoinsInfoList(pageInfo)
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

// GetCoinsPublic 不需要鉴权的币接口
// @Tags Coins
// @Summary 不需要鉴权的币接口
// @accept application/json
// @Produce application/json
// @Param data query coinReq.CoinsSearch true "分页获取币列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /coins/getCoinsPublic [get]
func (coinsApi *CoinsApi) GetCoinsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的币接口信息",
	}, "获取成功", c)
}
