package cryptocompare

import (
	"fmt"
	"strings"

	"github.com/blcokchina110/coinprice/xhttp"
	"github.com/blcokchina110/coinprice/xtime"

	"github.com/blcokchina110/coinprice/common"

	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
)

const (
	apiurl = "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%v&tsyms=USD"
)

type CryptoCompare struct {
	coinName  string
	timestamp int64
}

type cryptoCompareInfo struct {
	Raw     map[string]interface{} `json:"RAW"`
	DISPLAY interface{}            `json:"DISPLAY"`
}

type usd struct {
	TYPE       string  `json:"TYPE"`
	MARKET     string  `json:"MARKET"`
	FROMSYMBOL string  `json:"FROMSYMBOL"`
	TOSYMBOL   string  `json:"TOSYMBOL"`
	FLAGS      string  `json:"FLAGS"`
	PRICE      float64 `json:"PRICE"`
	LASTUPDATE int64   `json:"LASTUPDATE"`
}

//
func NewCryptoCompare(coinName string) *CryptoCompare {
	return &CryptoCompare{
		coinName:  strings.ToUpper(coinName),
		timestamp: xtime.Second(),
	}
}

//接口渠道
func (e *CryptoCompare) Name() string {
	return "cryptocompare"
}

//币种名称
func (e *CryptoCompare) CoinName() string {
	return e.coinName
}

//交易对
func (e *CryptoCompare) Pair() string {
	return e.coinName + "/" + common.UpperUSD
}

//时间戳
func (e *CryptoCompare) TimeStamp() int64 {
	return e.timestamp
}

//获取指定币种美元价格
func (e *CryptoCompare) GetPrice() decimal.Decimal {
	var info *cryptoCompareInfo
	if err := xhttp.GetDataUnmarshal(fmt.Sprintf(apiurl, e.coinName), nil, &info); err != nil {
		return decimal.NewFromInt(0)
	}

	usdInfo := info.Raw[e.coinName].(map[string]interface{})[common.UpperUSD]
	var result usd
	if err := mapstructure.Decode(usdInfo, &result); err == nil {
		e.timestamp = result.LASTUPDATE
		if result.FROMSYMBOL == e.coinName && xtime.CheckTimeValid(result.LASTUPDATE, 2) {
			return decimal.NewFromFloat(result.PRICE).Truncate(2)
		}
	}

	return decimal.NewFromInt(0)
}
