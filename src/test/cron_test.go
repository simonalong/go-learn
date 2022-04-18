package test

import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
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
