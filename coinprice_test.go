package coinprice

import (
	"github.com/blcokchina110/coinprice/cryptocompare"
	"fmt"
	"testing"
)

func TestCoinPrice(t *testing.T) {

	coinPriceEvent := NewCoinPriceEvent(cryptocompare.NewCryptoCompare("ltc"))
	fmt.Println(coinPriceEvent.CoinPrice.GetPrice())
}
