package main

import (
	"fmt"
	"time"
)

func main() {

	//now := time.Now()
	//// 格式化
	//fmt.Println(now.Format("2006年01月02 15-04.05.23123"))
	//
	//// 1630809826
	//// 秒
	//fmt.Println(now.Unix())
	//// 毫秒
	//fmt.Println(now.UnixNano() / 1e6)
	//// 纳秒
	//fmt.Println(now.UnixNano())
	//

	//字符串类型转time
	//s4 := "1999年10月19日" //字符串
	t4, err := time.ParseInLocation("2006-01-02 15:04:05.000", "2018-09-10 00:00:00.124", time.Local)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(t4)
	//
	//data := time.Time{}
	//fmt.Println(data)
	//fmt.Println(data.Unix())
	//
	//emptyTime := time.Time{}
	//
	//if data == emptyTime {
	//	fmt.Println("空")
	//}
}

func AddSecond(times time.Time, second uint) time.Time {
	var du time.Duration = 500 * time.Millisecond
	return times.Add(du)
}
