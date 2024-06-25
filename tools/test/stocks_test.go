package test

import (
	"fmt"
	baseTime "github.com/isyscore/isc-gobase/time"
	"math/rand"
	"testing"
	"time"
)

var timeForm = "20060102"

//var timeForm = "2006年01月02日"

var festivals = []string{
	// 元旦
	"20220103",
	// 春节
	"20220131",
	"20220201",
	"20220202",
	"20220203",
	"20220204",
	// 清明节
	"20220404",
	// 劳动节
	"20220502",
	"20220503",
	"20220504",
	// 端午节
	"20220603",
	// 中秋节
	"20220912",
	// 国庆节
	"20221003",
	"20221004",
	"20221005",
	"20221006",
	"20221007",
}

// 是否是休息日
func isRest(timeDay time.Time) bool {
	// 节日
	timeStr := baseTime.TimeToStringFormat(timeDay, "20060102")
	for _, festival := range festivals {
		if festival == timeStr {
			return true
		}
	}

	// 周六周日
	week := timeDay.Weekday()
	if week.String() == "Saturday" || week.String() == "Sunday" {
		return true
	}
	return false
}

// 获取不是节假日的昨天
func getYesterdayWithoutRest(yesterdayFirst time.Time) time.Time {
	// 如果是休息日，向前推一天
	if isRest(yesterdayFirst) {
		return getYesterdayWithoutRest(yesterdayFirst.AddDate(0, 0, -1))
	}
	return yesterdayFirst
}

// 打印竞价的条件
func TestPrint(t *testing.T) {
	//策略1：formt := "昨日涨停；昨日连续涨停次数；昨日涨停原因；昨日涨停价成交额；今日竞价未涨停；今日竞价涨幅>=0；今日竞价额>昨日成交额*0.15；市值<100亿；今日竞价额>2000万；非次新；非ST；主板；"
	//策略2：formt := "20230105涨停；20230105连续涨停次数；20230105涨停原因；20230105涨停价成交额；20230106竞价未涨停；20230106竞价涨幅>=0；20230106竞价额>20230105成交额*0.1；市值<200亿；20230106竞价额>2000万；非ST；主板；"
	//策略2：formt := "20230105涨停；20230105连续涨停次数；20230105涨停原因；20230105涨停价成交额；20230106竞价未涨停；20230106竞价涨幅>=0；20230106竞价额>20230105成交额*0.15；市值<200亿；20230106竞价额>3000万；非ST；主板；"
	formt := "%s涨停；%s连续涨停次数；%s涨停原因；%s涨停价成交额；%s竞价未涨停；%s竞价涨幅>=0；%s竞价额>%s成交额*0.1；%s竞价额>1000万；市值<200亿；非ST；主板；所属行业"

	//nowDay := time.Now()
	nowDay := baseTime.ParseTime("2022-01-17")
	today := nowDay
	lastMonth := today.Month()
	for i := 0; i < 800; i++ {
		today = today.AddDate(0, 0, 1)

		if isRest(today) {
			continue
		}

		yesterday := getYesterdayWithoutRest(today.AddDate(0, 0, -1))
		yesStr := baseTime.TimeToStringFormat(yesterday, timeForm)
		todStr := baseTime.TimeToStringFormat(today, timeForm)

		if today.Month() != lastMonth {
			lastMonth = today.Month()
			fmt.Println()
		}

		fmt.Println(fmt.Sprintf(formt, yesStr, yesStr, yesStr, yesStr, todStr, todStr, todStr, yesStr, todStr))
	}
}

// 打印竞价的条件
func TestPrintTime2(t *testing.T) {
	//nowDay := time.Now()
	nowDay := baseTime.ParseTime("2022-01-01")
	lastMonth := nowDay.Month()
	for i := 1; i < 800; i++ {
		today := nowDay.AddDate(0, 0, i)
		if isRest(today) {
			continue
		}

		if today.Month() != lastMonth {
			lastMonth = today.Month()
			fmt.Println()
		}
		fmt.Println(baseTime.TimeToStringFormat(today, "2006年01月02日"))
	}
}

func TestRand(t *testing.T) {
	a := 1.2
	b := 2.5
	fmt.Println(rand.Float64()*(b-a) + a)
	fmt.Println(rand.Float64()*(b-a) + a)
	fmt.Println(rand.Float64()*(b-a) + a)
	fmt.Println(rand.Float64()*(b-a) + a)
	fmt.Println(rand.Float64()*(b-a) + a)
	fmt.Println(rand.Float64()*(b-a) + a)
}
