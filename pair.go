package coinprice

import (
	"strings"
)

//
type CurrencyPair struct {
	pairName    string
	currency1   string
	currency2   string
	description string
}

//
func NewCurrencyPair(currency1, currency2 string) *CurrencyPair {
	pair := new(CurrencyPair)
	if currency1 != "" && currency2 != "" {
		pair.currency1 = strings.ToUpper(currency1)
		pair.currency2 = strings.ToUpper(currency2)
		pair.pairName = pair.currency1 + "-" + pair.currency2
	}
	return pair
}

//交易对名称
func (c *CurrencyPair) PairName() string {
	return c.pairName
}

//是否反转
func (c *CurrencyPair) Reverse() bool {
	return c.check()
}

func (c *CurrencyPair) Currency1() string {
	if c.check() {
		return c.currency2
	}
	return c.currency1
}

func (c *CurrencyPair) Currency2() string {
	if c.check() {
		return c.currency1
	}
	return c.currency2
}

//检查交易对是否包含USD开头
func (c *CurrencyPair) check() bool {
	return strings.Contains(c.currency1, upperUSD) &&
		!strings.Contains(c.currency2, upperUSD)
}
