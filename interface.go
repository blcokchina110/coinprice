package coinprice

import (
	"github.com/shopspring/decimal"
)

//
type CoinPrice interface {
	//获取以太坊的美元价格
	GetPrice() decimal.Decimal
	//接口渠道
	Name() string
	//时间戳
	TimeStamp() int64
}
