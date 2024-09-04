// 自动生成模板Coins
package coin

import (
	"time"

	"github.com/Grace1China/cointown/server/global"
)

// 币 结构体  Coins
type Coins struct {
	global.GVA_MODEL
	// id              *time.Time `json:"id" form:"id" gorm:"uniqueIndex;column:id;comment:;" binding:"required"`                  //id
	TokenName       string    `json:"tokenName" form:"tokenName" gorm:"column:token_name;comment:;size:100;"`                  //币名
	Symbal          string    `json:"symbal" form:"symbal" gorm:"column:symbal;comment:;size:50;"`                             //币名
	ContractAddress string    `json:"contractAddress" form:"contractAddress" gorm:"column:contract_address;comment:;size:50;"` //合约地址
	BondingCurve    string    `json:"bondingCurve" form:"bondingCurve" gorm:"column:bonding_curve;comment:;size:50;"`          //Bond曲线
	Progress        string    `json:"progress" form:"progress" gorm:"column:progress;comment:;size:50;"`                       //progress
	Balance         string    `json:"balance" form:"balance" gorm:"column:balance;comment:;size:50;"`                          //余额
	Price           float64   `json:"price" form:"price" gorm:"column:price;comment:;size:50;"`                                //价格
	CreatedTime     string    `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;size:50;"`             //创建时间
	MarketCap       string    `json:"marketCap" form:"marketCap" gorm:"column:market_cap;comment:;size:50;"`                   //市场容量
	CommentsCount   string    `json:"commentsCount" form:"commentsCount" gorm:"column:comments_count;comment:;size:50;"`       //评价数量
	DevWallet       string    `json:"devWallet" form:"devWallet" gorm:"column:dev_wallet;comment:;size:50;"`                   //开发钱包
	DevHolding      string    `json:"devHolding" form:"devHolding" gorm:"column:dev_holding;comment:;size:50;"`                //开发钱包
	TwitterLink     string    `json:"twitterLink" form:"twitterLink" gorm:"column:twitter_link;comment:;size:50;"`             //X链接
	TelegramLink    string    `json:"telegramLink" form:"telegramLink" gorm:"column:telegram_link;comment:;size:50;"`          //T链接
	WebsiteLink     string    `json:"websiteLink" form:"websiteLink" gorm:"column:website_link;comment:;size:50;"`             //Web链接
	LivePriceLink   string    `json:"livePriceLink" form:"livePriceLink" gorm:"column:live_price_link;comment:;size:50;"`      //Price链接
	PriceWave5m     string    `json:"priceWave5m" form:"priceWave5m" gorm:"column:price__wave_5m;comment:;size:50;"`           //PriceWave5m
	PriceWave1h     string    `json:"priceWave1h" form:"priceWave1h" gorm:"column:price_wave_1h;comment:;size:50;"`            //PriceWave1h
	PriceWave6h     string    `json:"priceWave6h" form:"priceWave6h" gorm:"column:price_wave_6h;comment:;size:50;"`            //PriceWave6h
	TradeVol1m      string    `json:"tradeVol1m" form:"tradeVol1m" gorm:"column:trade_vol_1m;comment:;size:50;"`               //PriceWave6h
	TradeVol5m      string    `json:"tradeVol5m" form:"tradeVol5m" gorm:"column:trade_vol_5m;comment:;size:50;"`               //PriceWave6h
	TradeVol1h      string    `json:"tradeVol1h" form:"tradeVol1h" gorm:"column:trade_vol_1h;comment:;size:50;"`               //PriceWave6h
	BuyVol1m        string    `json:"buyVol1m" form:"buyVol1m" gorm:"column:buy_vol_1m;comment:;size:50;"`                     //PriceWave6h
	BuyVol5m        string    `json:"buyVol5m" form:"buyVol5m" gorm:"column:buy_Vol_5m;comment:;size:50;"`                     //PriceWave6h
	BuyVol1h        string    `json:"buyVol1h" form:"buyVol1h" gorm:"column:buy_Vol_1h;comment:;size:50;"`                     //PriceWave6h
	SaleVol1m       string    `json:"saleVol1m" form:"saleVol1m" gorm:"column:sale_vol_1m;comment:;size:50;"`                  //PriceWave6h
	SaleVol5m       string    `json:"saleVol5m" form:"saleVol5m" gorm:"column:sale_Vol_5m;comment:;size:50;"`                  //PriceWave6h
	SaleVol1h       string    `json:"saleVol1h" form:"saleVol1h" gorm:"column:sale_Vol_1h;comment:;size:50;"`                  //PriceWave6h
	TXs             string    `json:"tXs" form:"tXs" gorm:"column:t_x_s;comment:;size:50;"`                                    //TXs
	Vol             string    `json:"vol" form:"vol" gorm:"column:vol;comment:;size:50;"`                                      //vol
	Liq             string    `json:"liq" form:"liq" gorm:"column:liq;comment:;size:50;"`                                      //Liq
	Holder          string    `json:"holder" form:"holder" gorm:"column:holder;comment:;size:50;"`                             //Holder
	Open            string    `json:"open" form:"open" gorm:"column:open;comment:;size:50;"`                                   //Open
	NoMint          string    `json:"noMint" form:"noMint" gorm:"column:no_mint;comment:;size:50;"`                            //NoMint
	BlackList       string    `json:"blackList" form:"blackList" gorm:"column:black_list;comment:;size:50;"`                   //BlackList
	Burnt           string    `json:"burnt" form:"burnt" gorm:"column:burnt;comment:;size:50;"`                                //Burnt
	Top10           string    `json:"top10" form:"top10" gorm:"column:top10;comment:;size:50;"`                                //Top10
	DevAddr         string    `json:"devAddr" form:"devAddr" gorm:"column:dev_addr;comment:;size:50;"`                         //DevAddr
	DevInfo         string    `json:"devInfo" form:"devInfo" gorm:"column:dev_info;comment:;size:50;"`                         //DevInfo
	DevBurn         string    `json:"devBurn" form:"devBurn" gorm:"column:dev_burn;comment:;size:2000;"`                       //DevBurn
	AdvInfo         string    `json:"advInfo" form:"advInfo" gorm:"column:adv_info;comment:;size:2000;"`                       //AdvInfo
	Circulating     string    `json:"circulating" form:"circulating" gorm:"column:circulating;comment:;size:50;"`              //AdvInfo
	CreatorBal      string    `json:"creatorBal" form:"creatorBal" gorm:"column:creator_bal;comment:;size:50;"`                //AdvInfo
	CreatorBalPer   string    `json:"creatorBalPer" form:"creatorBalPer" gorm:"column:creator_bal_per;comment:;size:50;"`      //AdvInfo
	MintCoins       string    `json:"mintCoins" form:"mintCoins" gorm:"column:mint_coins;comment:;size:50;"`                   //AdvInfo
	RecommendAddr   string    `json:"recommendAddr" form:"recommendAddr" gorm:"column:recommend_addr;comment:;size:50;"`       //AdvInfo
	MsgId           uint      `json:"msgId" form:"msgId" gorm:"column:msg_id;comment:;size:50;"`                               //AdvInfo
	PriceTest1      float64   `json:"priceTest1" form:"priceTest1" gorm:"column:price_test_1;comment:;size:50;"`
	PriceTest2      float64   `json:"priceTest2" form:"priceTest2" gorm:"column:price_test_2;comment:;size:50;"`
	PriceTest3      float64   `json:"priceTest3" form:"priceTest3" gorm:"column:price_test_3;comment:;size:50;"`
	PriceTestWave1  float64   `json:"priceTestWave1" form:"priceTestWave1" gorm:"column:price_test_wave_1;comment:;size:50;"`
	PriceTestWave2  float64   `json:"priceTestWave2" form:"priceTestWave2" gorm:"column:price_test_wave_2;comment:;size:50;"`
	PriceTestWave3  float64   `json:"priceTestWave3" form:"priceTestWave3" gorm:"column:price_test_wave_3;comment:;size:50;"`
	PriceTestTime1  time.Time `json:"priceTestTime1" form:"priceTestTime1" gorm:"column:price_test_time_1;comment:;"`
	PriceTestTime2  time.Time `json:"priceTestTime2" form:"priceTestTime2" gorm:"column:price_test_time_2;comment:;"`
	PriceTestTime3  time.Time `json:"priceTestTime3" form:"priceTestTime3" gorm:"column:price_test_time_3;comment:;"`
	SmartBuyer      uint      `json:"smartBuyer" form:"smartBuyer" gorm:"column:smart_buyer;comment:;"`
	ChatId          string    `json:"chatId" form:"chatId" gorm:"column:chat_id;comment:;size:50;"` //AdvInfo

}

// TableName 币 Coins自定义表名 coins
func (Coins) TableName() string {
	return "coins"
}
