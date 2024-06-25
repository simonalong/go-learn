package test

import (
	"fmt"
	"testing"

	"github.com/robfig/cron"
)

func TestCron1(t *testing.T) {
	c := cron.New() //精确到秒

	// 5分钟执行一次
	spec := "0 0/5 * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println("11111")
	})

	c.Start()
	select {} //阻塞主线程停止
}

func TestFloatToInt(t *testing.T) {

	var data float64
	data = 1663211939089

	var tem = int64(data)
	fmt.Println(tem)
}
