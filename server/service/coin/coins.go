package coin

import (
	"fmt"
	"time"

	"github.com/Grace1China/cointown/server/global"
	"github.com/Grace1China/cointown/server/model/coin"
	coinReq "github.com/Grace1China/cointown/server/model/coin/request"
)

type CoinsService struct{}

// CreateCoins 创建币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) CreateCoins(coins *coin.Coins) (err error) {
	err = global.GVA_DB.Create(coins).Error
	return err
}

// DeleteCoins 删除币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) DeleteCoins(ID string) (err error) {
	err = global.GVA_DB.Delete(&coin.Coins{}, "id = ?", ID).Error
	return err
}

// DeleteCoinsByIds 批量删除币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) DeleteCoinsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]coin.Coins{}, "id in ?", IDs).Error
	return err
}

// UpdateCoins 更新币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) UpdateCoins(coins coin.Coins) (err error) {
	err = global.GVA_DB.Model(&coin.Coins{}).Where("id = ?", coins.ID).Updates(&coins).Error
	return err
}

// GetCoins 根据ID获取币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) GetCoins(ID string) (coins coin.Coins, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&coins).Error
	return
}

// GetCoins 根据ID获取币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) GetCoinsByAddr(contract_address string) (coins coin.Coins, err error) {
	err = global.GVA_DB.Where("contract_address = ?", contract_address).First(&coins).Error
	return
}

// GetCoins 根据ID获取币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) GetCoinsByTimeDue(t time.Time) (list []coin.Coins, err error) {
	// 创建db
	db := global.GVA_DB.Model(&coin.Coins{})
	var coinss []coin.Coins
	db = db.Where("((price_test_time_1 < ? AND price_test_1<=0 AND price_test_1>-100) OR (price_test_time_2 < ? AND price_test_2>-100 AND price_test_2<=0 ) OR (price_test_time_3 < ? AND price_test_3>-100 AND price_test_3<=0))", t, t, t)
	fmt.Println("get coins =============================>", coinss)
	err = db.Find(&coinss).Error
	return coinss, err
}

// GetCoinsInfoList 分页获取币记录
// Author [piexlmax](https://github.com/piexlmax)
func (coinsService *CoinsService) GetCoinsInfoList(info coinReq.CoinsSearch) (list []coin.Coins, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&coin.Coins{})
	var coinss []coin.Coins
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id DESC")
	}

	err = db.Find(&coinss).Error
	return coinss, total, err
}
