package test

import (
	"fmt"
	baseTime "github.com/isyscore/isc-gobase/time"
	"testing"
	"time"
)

func TestP(t *testing.T) {
	var a int = 10
	//故意让10除以0
	var b int
	//recover应当放在可能出现的错误之前
	//通俗理解为宜未雨绸缪，不可亡羊补牢
	defer func() {
		//放在匿名函数里,err捕获到错误信息，并且输出
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	c := a / b
	fmt.Println(c)
}

// 股票生成
func TestPrint(t *testing.T) {
	formt := "%s首板涨停；%s竞价开盘涨幅；%s连涨天数；%s涨停价成交额；%s竞价额>%s成交额*0.15；非ST；主板；%s竞价未涨停；市值小于200亿；"
	timeForm := "2006年01月02日"
	nowDay := time.Now()
	for i := 0; i < 60; i++ {
		tod := nowDay.AddDate(0, 0, -i)
		week := tod.Weekday()
		var yes time.Time
		if week.String() == "Monday" {
			yes = nowDay.AddDate(0, 0, -3)
		} else if week.String() == "Saturday" || week.String() == "Sunday" {
			continue
		} else {
			yes = tod.AddDate(0, 0, -1)
		}

		yesStr := baseTime.TimeToStringFormat(yes, timeForm)
		todStr := baseTime.TimeToStringFormat(tod, timeForm)
		fmt.Println(fmt.Sprintf(formt, yesStr, todStr, yesStr, yesStr, todStr, yesStr, todStr))
	}
}
