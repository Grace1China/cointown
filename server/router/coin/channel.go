package coin

import (
	"github.com/Grace1China/cointown/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChannelRouter struct{}

// InitChannelRouter 初始化 频道 路由信息
func (s *ChannelRouter) InitChannelRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	channelRouter := Router.Group("channel").Use(middleware.OperationRecord())
	channelRouterWithoutRecord := Router.Group("channel")
	channelRouterWithoutAuth := PublicRouter.Group("channel")
	{
		// channelRouter.POST("createChannel", channelApi.CreateChannel)   // 新建频道
		channelRouter.DELETE("deleteChannel", channelApi.DeleteChannel)           // 删除频道
		channelRouter.DELETE("deleteChannelByIds", channelApi.DeleteChannelByIds) // 批量删除频道
		channelRouter.PUT("updateChannel", channelApi.UpdateChannel)              // 更新频道
	}
	{
		channelRouterWithoutRecord.GET("findChannel", channelApi.FindChannel)       // 根据ID获取频道
		channelRouterWithoutRecord.GET("getChannelList", channelApi.GetChannelList) // 获取频道列表
	}
	{
		channelRouterWithoutAuth.POST("createChannel", channelApi.CreateChannel)      // 新建频道
		channelRouterWithoutAuth.GET("getChannelPublic", channelApi.GetChannelPublic) // 获取频道列表
	}
}
