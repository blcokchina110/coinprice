package coinprice

import (
	"fmt"
	"testing"

	"github.com/blcokchina110/coinprice/currencypair"
	"github.com/blcokchina110/coinprice/ethereumdb"
)

func TestCoinPrice(t *testing.T) {
	c := currencypair.NewCurrencyPair("ETH", "USD")
	coinPriceEvent := NewCoinPriceEvent(ethereumdb.NewEthereumDB(c))
	fmt.Println(coinPriceEvent.CoinPrice.GetPrice())
}
