package cryptocompare

import (
	"fmt"

	"github.com/blcokchina110/coinprice/currencypair"
	"github.com/blcokchina110/coinprice/xhttp"
	"github.com/blcokchina110/coinprice/xtime"

	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
)

const (
	apiurl = "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%v&tsyms=%v"
)

type CryptoCompare struct {
	currencypair *currencypair.CurrencyPair
	timestamp    int64
}

type cryptoCompareInfo struct {
	Raw     map[string]interface{} `json:"RAW"`
	Display interface{}            `json:"DISPLAY"`
}

type coinPriceInfo struct {
	Type       string  `json:"TYPE"`
	Market     string  `json:"MARKET"`
	FromSymbol string  `json:"FROMSYMBOL"`
	ToSymbol   string  `json:"TOSYMBOL"`
	Flags      string  `json:"FLAGS"`
	Price      float64 `json:"PRICE"`
	LastUpdate int64   `json:"LASTUPDATE"`
}

//
func NewCryptoCompare(currencypair *currencypair.CurrencyPair) *CryptoCompare {
	return &CryptoCompare{
		currencypair: currencypair,
		timestamp:    xtime.Second(),
	}
}

//接口渠道
func (e *CryptoCompare) Name() string {
	return "cryptocompare"
}

//时间戳
func (e *CryptoCompare) TimeStamp() int64 {
	return e.timestamp
}

//获取指定币种美元价格
func (e *CryptoCompare) GetPrice() decimal.Decimal {
	var info *cryptoCompareInfo
	if err := xhttp.GetDataUnmarshal(fmt.Sprintf(apiurl, e.currencypair.Currency1(), e.currencypair.Currency2()), nil, &info); err != nil {
		return decimal.NewFromInt(0)
	}

	priceInfo := info.Raw[e.currencypair.Currency1()].(map[string]interface{})[e.currencypair.Currency2()]
	var result coinPriceInfo
	if err := mapstructure.Decode(priceInfo, &result); err == nil {
		e.timestamp = result.LastUpdate
		if result.FromSymbol == e.currencypair.Currency1() && xtime.CheckTimeValid(result.LastUpdate, 2) {
			df := decimal.NewFromFloat(result.Price)
			if !e.currencypair.Reverse() {
				return df.Truncate(2)
			}
			return decimal.NewFromInt(1).Div(df).Truncate(8)
		}
	}

	return decimal.NewFromInt(0)
}
