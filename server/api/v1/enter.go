package v1

import (
	"github.com/Grace1China/cointown/server/api/v1/coin"
	"github.com/Grace1China/cointown/server/api/v1/example"
	"github.com/Grace1China/cointown/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CoinApiGroup    coin.ApiGroup
}
