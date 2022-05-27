package coinprice

import (
	"fmt"

	"github.com/blcokchina110/coinprice/xtime"
	"github.com/blcokchina110/xhttp"

	"github.com/shopspring/decimal"
)

const ()

type EthereumDB struct {
	currencypair *CurrencyPair
	timestamp    int64
}

//
type ethereumDBInfo struct {
	Price          decimal.Decimal `json:"price"`
	QuoteVolume24h decimal.Decimal `json:"quoteVolume24h"`
	Timestamp      int64           `json:"timestamp"`
}

//
func NewEthereumDB(currencypair *CurrencyPair) *EthereumDB {
	return &EthereumDB{
		currencypair: currencypair,
		timestamp:    xtime.Second(),
	}
}

//接口渠道
func (e *EthereumDB) Name() string {
	return "ethereumdb"
}

//时间戳
func (e *EthereumDB) TimeStamp() int64 {
	return e.timestamp
}

//获取指定币种美元价格
func (e *EthereumDB) GetPrice() decimal.Decimal {
	var infos []ethereumDBInfo

	url := fmt.Sprintf(ethereumdbApiUrl, e.currencypair.Currency1(), e.currencypair.Currency2())
	if err := xhttp.GetParseData(url, nil, &infos); err == nil {
		if len(infos) > 0 {
			e.timestamp = infos[0].Timestamp
			if !e.currencypair.Reverse() {
				return infos[0].Price.Truncate(2)
			}
			return decimal.NewFromInt(1).Div(infos[0].Price).Truncate(8)
		}
	}

	return decimal.NewFromInt(0)
}
