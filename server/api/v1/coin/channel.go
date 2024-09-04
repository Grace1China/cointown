package coin

import (
	"errors"

	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"
	coinReq "github.com/Grace1China/cointown/server/model/coin/request"
	"github.com/Grace1China/cointown/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	// "bytes"
	// "encoding/json"
	// "net/http"
)

type ChannelApi struct{}

func (channelApi *ChannelApi) CreateUpdtChannel(c *gin.Context) {
	ID := c.Query("ChatId")
	rechannel, err := channelService.GetChannel(ID)
	if err != nil {
		//创建新的频道
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var channel coin.Channel
			err := c.ShouldBindJSON(&channel)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}

			err = channelService.CreateChannel(&channel)
			if err != nil {
				global.GVA_LOG.Error("创建失败!", zap.Error(err))
				response.FailWithMessage("创建失败:"+err.Error(), c)
				return
			}
			response.OkWithMessage("创建成功", c)
			return
		} else {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else {
		//更新频道
		rechannel.LastMessageText = c.Param("LastMessageText")
		err = channelService.UpdateChannel(rechannel)
		if err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败:"+err.Error(), c)
			return
		}
		response.OkWithMessage("更新成功", c)
		return
	}
}

// CreateChannel 创建频道
// @Tags Channel
// @Summary 创建频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.Channel true "创建频道"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /channel/createChannel [post]
func (channelApi *ChannelApi) CreateChannel(c *gin.Context) {
	var channel coin.Channel
	err := c.ShouldBindJSON(&channel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = channelService.CreateChannel(&channel)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	// 创建成功，返回对象和消息
	response.OkWithData(channel, c)
}

// DeleteChannel 删除频道
// @Tags Channel
// @Summary 删除频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.Channel true "删除频道"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /channel/deleteChannel [delete]
func (channelApi *ChannelApi) DeleteChannel(c *gin.Context) {
	ID := c.Query("ID")
	err := channelService.DeleteChannel(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteChannelByIds 批量删除频道
// @Tags Channel
// @Summary 批量删除频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /channel/deleteChannelByIds [delete]
func (channelApi *ChannelApi) DeleteChannelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := channelService.DeleteChannelByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateChannel 更新频道
// @Tags Channel
// @Summary 更新频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body coin.Channel true "更新频道"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /channel/updateChannel [put]
func (channelApi *ChannelApi) UpdateChannel(c *gin.Context) {
	var channel coin.Channel
	err := c.ShouldBindJSON(&channel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = channelService.UpdateChannel(channel)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindChannel 用id查询频道
// @Tags Channel
// @Summary 用id查询频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query coin.Channel true "用id查询频道"
// @Success 200 {object} response.Response{data=coin.Channel,msg=string} "查询成功"
// @Router /channel/findChannel [get]
func (channelApi *ChannelApi) FindChannel(c *gin.Context) {
	ID := c.Query("ID")
	rechannel, err := channelService.GetChannel(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rechannel, c)
}

// GetChannelList 分页获取频道列表
// @Tags Channel
// @Summary 分页获取频道列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query coinReq.ChannelSearch true "分页获取频道列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /channel/getChannelList [get]
func (channelApi *ChannelApi) GetChannelList(c *gin.Context) {
	var pageInfo coinReq.ChannelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := channelService.GetChannelInfoList(pageInfo)
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

// GetChannelPublic 不需要鉴权的频道接口
// @Tags Channel
// @Summary 不需要鉴权的频道接口
// @accept application/json
// @Produce application/json
// @Param data query coinReq.ChannelSearch true "分页获取频道列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /channel/getChannelPublic [get]
func (channelApi *ChannelApi) GetChannelPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的频道接口信息",
	}, "获取成功", c)
}
