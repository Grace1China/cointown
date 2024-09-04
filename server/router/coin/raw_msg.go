package coin

import (
	"github.com/Grace1China/cointown/server/middleware"
	"github.com/gin-gonic/gin"
)

type RawMsgRouter struct{}

// InitRawMsgRouter 初始化 原始消息 路由信息
func (s *RawMsgRouter) InitRawMsgRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	rawmsgRouter := Router.Group("rawmsg").Use(middleware.OperationRecord())
	rawmsgRouterWithoutRecord := Router.Group("rawmsg")
	rawmsgRouterWithoutAuth := PublicRouter.Group("rawmsg")
	{
		// rawmsgRouter.POST("createRawMsg", rawmsgApi.CreateRawMsg)   // 新建原始消息
		rawmsgRouter.DELETE("deleteRawMsg", rawmsgApi.DeleteRawMsg)                 // 删除原始消息
		rawmsgRouter.DELETE("deleteRawMsgByIds", rawmsgApi.DeleteRawMsgByIds)       // 批量删除原始消息
		rawmsgRouter.PUT("updateRawMsg", rawmsgApi.UpdateRawMsg)                    // 更新原始消息
		rawmsgRouter.POST("deleteRawMsg7D", rawmsgApi.DeleteRawMsg7D)               // 更新原始消息
		rawmsgRouter.POST("UpdateRawMsgStatusAll", rawmsgApi.UpdateRawMsgStatusAll) // 更新原始消息

	}
	{
		rawmsgRouterWithoutRecord.GET("findRawMsg", rawmsgApi.FindRawMsg) // 根据ID获取原始消息
		// rawmsgRouterWithoutRecord.GET("getRawMsgList", rawmsgApi.GetRawMsgList) // 获取原始消息列表
	}
	{
		rawmsgRouterWithoutAuth.GET("getRawMsgPublic", rawmsgApi.GetRawMsgPublic) // 获取原始消息列表
		rawmsgRouterWithoutAuth.POST("createRawMsg", rawmsgApi.CreateRawMsg)      // 新建原始消息
		rawmsgRouterWithoutAuth.GET("getRawMsgList", rawmsgApi.GetRawMsgList)     // 获取原始消息列表

	}
}
