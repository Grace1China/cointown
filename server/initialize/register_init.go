package initialize

import (
	_ "github.com/Grace1China/cointown/server/source/example"
	_ "github.com/Grace1China/cointown/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
