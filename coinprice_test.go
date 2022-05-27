package coinprice

import (
	"fmt"
	"testing"
)

func TestCoinPrice(t *testing.T) {
	c := NewCurrencyPair("ETH", "USD")
	coinPriceEvent := NewCoinPriceEvent(NewEthereumDB(c))
	fmt.Println(coinPriceEvent.CoinPrice.GetPrice())
}
