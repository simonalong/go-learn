package main

import (
	"fmt"
	"time"
)

var (
	YEAR  = "2006"
	MONTH = "01"
	DAY   = "02"

	HOUR   = "15"
	MINUTE = "04"
	SECOND = "05"

	yMdHmsSSS = "2006-01-02 15:04:05.000"
	yMdHmsS   = "2006-01-02 15:04:05.0"
	yMdHms    = "2006-01-02 15:04:05"
	yMdHm     = "2006-01-02 15:04"
	yMdH      = "2006-01-02 15"
	yMd       = "2006-01-02"
	yM        = "2006-01"
	y         = "2006"
	yyyyMMdd  = "20060102"

	HmsSSSMore = "15:04:05.SSSSSSSSS"
	HmsSSS     = "15:04:05.000"
	Hms        = "15:04:05"
	Hm         = "15:04"
	H          = "15"
)

//1. time -> string
//2. string -> time
//3. time -> long
//4. long -> time
//5. 时区
func main() {

	now := time.Now()
	// 格式化
	// 2021年12月24 11-37.03.1723
	fmt.Println(now.Format("2006年01月02 15-04.05.0000"))

	// 年 2021
	fmt.Println(now.Year())
	// 月 12
	fmt.Println(now.Month())
	// 天 24
	fmt.Println(now.Day())
	// 时 11
	fmt.Println(now.Hour())
	// 分 37
	fmt.Println(now.Minute())
	// 秒 03
	fmt.Println(now.Second())
	// 毫秒 172332000
	fmt.Println(now.Nanosecond())

	// 转化对应的单位
	// 秒 1640325398
	fmt.Println(now.Unix())
	// 毫秒 1640325398551
	fmt.Println(now.UnixMilli())
	// 微秒 1640325398551894
	fmt.Println(now.UnixMicro())
	// 纳秒 1640325398551894000
	fmt.Println(now.UnixNano())

	// 东八区
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	fmt.Println(time.Now().In(cstZone).Format("2006-01-02 15:04:05"))

	// 北京上海时间
	l, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(timeToStringYmdHms(time.Now().In(l)))

	// 美国时间
	l, _ = time.LoadLocation("America/Adak")
	fmt.Println(timeToStringYmdHms(time.Now().In(l)))

	t2, _ := parseTimeYmsHmsS("2012-11-01 22:08:41.321")
	fmt.Println(timeToStringYmdHmsS(t2))

	t3, _ := parseTimeYmsHmsS("2012-11-01 22:08:41.123")
	fmt.Println(timeToStringYmdHmsS(t3))
}

func timeToStringYmdHms(t time.Time) string {
	return t.Format(yMdHms)
}

func timeToStringYmdHmsS(t time.Time) string {
	return t.Format(yMdHmsSSS)
}

func timeToStringFormat(t time.Time, format string) string {
	return t.Format(format)
}

func parseTimeYmsHms(timeStr string) (time.Time, error) {
	return time.ParseInLocation(YEAR+"-"+MONTH+"-"+DAY+" "+HOUR+":"+MINUTE+":"+SECOND, timeStr, time.Local)
}

func parseTimeYmsHmsS(timeStr string) (time.Time, error) {
	return time.ParseInLocation(YEAR+"-"+MONTH+"-"+DAY+" "+HOUR+":"+MINUTE+":"+SECOND+".000", timeStr, time.Local)
}

func parseTimeYmsHmsLoc(timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(YEAR+"-"+MONTH+"-"+DAY+" "+HOUR+":"+MINUTE+":"+SECOND, timeStr, loc)
}

func parseTimeYmsHmsSLoc(timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(YEAR+"-"+MONTH+"-"+DAY+" "+HOUR+":"+MINUTE+":"+SECOND+".000", timeStr, loc)
}
