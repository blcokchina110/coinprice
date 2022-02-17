package coinprice

import "github.com/shopspring/decimal"

//
type CoinPrice interface {
	//获取以太坊的美元价格
	GetPrice() decimal.Decimal
	//接口渠道
	Name() string
	//交易对
	Pair() string
	//币种名称
	CoinName() string
}
