package coin

import "github.com/Grace1China/cointown/server/service"

type ApiGroup struct {
	RawMsgApi
	CoinsApi
	ChannelApi
}

var (
	rawmsgService    = service.ServiceGroupApp.CoinServiceGroup.RawMsgService
	coinsService     = service.ServiceGroupApp.CoinServiceGroup.CoinsService
	channelService   = service.ServiceGroupApp.CoinServiceGroup.ChannelService
	tempParseService = service.ServiceGroupApp.CoinServiceGroup.TempParseService
)
