package xtime

import (
	"fmt"
	"strconv"
	"time"
)

const (
	goCreateTime = "2006-01-02 15:04:05"
	goDate       = "2006-01-02"
	goTime       = "15:04:05"
)

//获取毫秒
func MilliSecond() int64 {
	return time.Now().UnixNano() / 1e6
}

//获取秒
func Second() int64 {
	return time.Now().Unix()
}

//
func AddSecondByDay(day int64) int64 {
	return Second() + day*96400
}

//
func GetDateTimeUnix(timeType int) int64 {
	switch timeType {
	case 1:
		return getDateUnix("00:00:00")
	case 2:
		return getDateUnix("23:59:59")
	}
	return 0
}

//至今天 23:59:59剩余秒数
func ToTodaySecond() int64 {
	now := time.Now().UTC()

	t := now.Format(goDate) + " 23:59:59"
	today, _ := time.Parse(goCreateTime, t)
	return today.Unix() - now.Unix() - 28800
}

//检查时间是否有效
func CheckTimeValid(second int64, minute int) bool {
	s := Second()
	if second >= s-int64(minute*60) && second <= s {
		return true
	}
	return false
}

//
func getDateUnix(t string) int64 {
	y, m, d := time.Now().Date()
	mm := func() string {
		nm := int(m)
		if nm < 10 {
			return fmt.Sprintf("0%v", nm)
		}
		return strconv.Itoa(nm)
	}()
	dt := fmt.Sprintf("%v-%v-%v %v", y, mm, d, t)

	p, _ := time.Parse(goCreateTime, dt)

	return p.UnixNano()/1e6 - 8*60*60*1000

}
