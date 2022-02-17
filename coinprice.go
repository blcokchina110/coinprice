package coinprice

//
type CoinPriceEvent struct {
	CoinPrice CoinPrice
}

//
func NewCoinPriceEvent(coinPrice CoinPrice) *CoinPriceEvent {
	return &CoinPriceEvent{
		CoinPrice: coinPrice,
	}
}
