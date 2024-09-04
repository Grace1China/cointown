package service

import (
	"github.com/Grace1China/cointown/server/service/coin"
	"github.com/Grace1China/cointown/server/service/example"
	"github.com/Grace1China/cointown/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CoinServiceGroup    coin.ServiceGroup
}
