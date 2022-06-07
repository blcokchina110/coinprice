package coinprice

import (
	"fmt"
	"testing"
)

func TestCoinPrice(t *testing.T) {
	c := NewCurrencyPair("ETH", "USD")
	coinPriceEvent := NewCoinPriceEvent(NewCryptoCompare(c))
	fmt.Println(coinPriceEvent.CoinPrice.GetPrice())
}
