package ethereumdb

import (
	"fmt"
	"strings"

	"github.com/blcokchina110/coinprice/common"
	"github.com/blcokchina110/coinprice/xhttp"
	"github.com/blcokchina110/coinprice/xtime"

	"github.com/shopspring/decimal"
)

const (
	apiurl = "https://api.ethereumdb.com/v1/timeseries?pair=%v-USD&range=10mi&type=line"
)

type EthereumDB struct {
	coinName  string
	timestamp int64
}

//
type ethereumDBInfo struct {
	Price          decimal.Decimal `json:"price"`
	QuoteVolume24h decimal.Decimal `json:"quoteVolume24h"`
	Timestamp      int64           `json:"timestamp"`
}

//
func NewEthereumDB(coinName string) *EthereumDB {
	return &EthereumDB{
		coinName:  strings.ToUpper(coinName),
		timestamp: xtime.Second(),
	}
}

//接口渠道
func (e *EthereumDB) Name() string {
	return "ethereumdb"
}

//币种名称
func (e *EthereumDB) CoinName() string {
	return e.coinName
}

//交易对
func (e *EthereumDB) Pair() string {
	return e.coinName + "/" + common.UpperUSD
}

//时间戳
func (e *EthereumDB) TimeStamp() int64 {
	return e.timestamp
}

//获取指定币种美元价格
func (e *EthereumDB) GetPrice() decimal.Decimal {
	var infos []ethereumDBInfo
	if err := xhttp.GetDataUnmarshal(fmt.Sprintf(apiurl, e.coinName), nil, &infos); err != nil {
		return decimal.NewFromInt(0)
	}

	if len(infos) > 0 {
		e.timestamp = infos[0].Timestamp
		return infos[0].Price.Truncate(2)
	}

	return decimal.NewFromInt(0)
}
