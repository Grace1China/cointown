package initialize

import (
	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(coin.RawMsg{}, coin.Coins{}, coin.Coins{}, coin.Channel{})
	if err != nil {
		return err
	}
	return nil
}
