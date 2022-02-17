package xtime

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	x := GetDateTimeUnix(2)
	fmt.Println(x)

	fmt.Println(CheckTimeValid(1645093654, 1))
}
