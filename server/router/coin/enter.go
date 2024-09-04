package coin

import api "github.com/Grace1China/cointown/server/api/v1"

type RouterGroup struct {
	RawMsgRouter
	CoinsRouter
	ChannelRouter
}

var (
	rawmsgApi  = api.ApiGroupApp.CoinApiGroup.RawMsgApi
	coinsApi   = api.ApiGroupApp.CoinApiGroup.CoinsApi
	channelApi = api.ApiGroupApp.CoinApiGroup.ChannelApi
)
