package test

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	var dat time.Duration = 12

	name := Name{
		// 正确
		//tt: 12 * time.Second,
		// 编译失败
		tt: dat * time.Second,
	}

	fmt.Println(name)
}

type Name struct {
	tt time.Duration
}
