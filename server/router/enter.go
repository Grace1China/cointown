package router

import (
	"github.com/Grace1China/cointown/server/router/coin"
	"github.com/Grace1China/cointown/server/router/example"
	"github.com/Grace1China/cointown/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Coin    coin.RouterGroup
}
