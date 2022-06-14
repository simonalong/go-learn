package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	start := time.Now()
	end := time.Now()

	fmt.Println(end.Second())
	fmt.Println(start.Second())
	fmt.Println(start.Nanosecond())
	fmt.Println(start.Hour())
	// 毫秒数
	fmt.Println(end.UnixMilli())
	// 微妙数
	fmt.Println(end.UnixMicro())
	// 纳秒数
	fmt.Println(end.UnixNano())
}
