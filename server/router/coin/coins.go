package coin

import (
	"github.com/Grace1China/cointown/server/middleware"
	"github.com/gin-gonic/gin"
)

type CoinsRouter struct{}

// InitCoinsRouter 初始化 币 路由信息
func (s *CoinsRouter) InitCoinsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	coinsRouter := Router.Group("coins").Use(middleware.OperationRecord())
	coinsRouterWithoutRecord := Router.Group("coins")
	coinsRouterWithoutAuth := PublicRouter.Group("coins")
	{
		// coinsRouter.POST("createCoins", coinsApi.CreateCoins)   // 新建币
		coinsRouter.DELETE("deleteCoins", coinsApi.DeleteCoins)           // 删除币
		coinsRouter.DELETE("deleteCoinsByIds", coinsApi.DeleteCoinsByIds) // 批量删除币
		coinsRouter.PUT("updateCoins", coinsApi.UpdateCoins)              // 更新币
	}
	{
		coinsRouterWithoutRecord.GET("findCoins", coinsApi.FindCoins)       // 根据ID获取币
		coinsRouterWithoutRecord.GET("getCoinsList", coinsApi.GetCoinsList) // 获取币列表
	}
	{
		coinsRouterWithoutAuth.POST("createCoins", coinsApi.CreateCoins)      // 新建币
		coinsRouterWithoutAuth.GET("getCoinsPublic", coinsApi.GetCoinsPublic) // 获取币列表
	}
}
