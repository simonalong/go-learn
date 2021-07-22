package main

import (
	"fmt"
	"go/types"
)

// 原来if可以不用写括号的唉
func main() {

	showIf(12)
	showIf(32)
	showIf(2)
	showIf2(43)

	readData()
	println(showSwitch(10))
	showSwitchType(10)
	showSwitchType(12.3)
	showSwitchType("")
	showSwitchMultiValue(2)
	showSelect1()
}

func showIf(a int) {

	if a == 12 {
		println("12")
	} else if a > 12 {
		println(" > 12")
	} else if a < 12 {
		println(" < 12")
	}
}

func showIf2(a int) {

	if a > 12 {
		if a > 23 {
			println("> 23")
		} else {
			println("<= 23")
		}
	}
}

func readData() {
	var a int
	b, _ := fmt.Scan(&a)
	if b > 23 {
		println("> 23")
	} else {
		println("<= 23")
	}
}

func showSwitch(a int) int {
	println("================================================")
	switch a {
	case 10:
		return 100
	case 20:
		return 200
	case 30:
		return 300
	default:
		return 10
	}
}

// 判断数据的类型
func showSwitchType(a interface{}) {
	switch a.(type) {
	case types.Nil:
		println("空")
	case int:
		println("int")
	case float32:
		println("float32")
	case float64:
		println("float64")
	case string:
		println("string")
	}
}

func showSwitchMultiValue(a int) {
	switch a {
	case 1, 2, 3, 4:
		println("1~4")
	case 6, 7:
		println("6,7")
	default:
		println("默认")
	}
}

// 下面这个select是跟并发结合使用的，后续
func showSelect1() {
	println("================================================")
	var c1, c2 chan int
	i1, i2 := 12, 32

	select {
	case i1 = <-c1:
		println("接收", i1, "from", c1)
	case c2 <- i2:
		println("向", c2, "数据", i2)
	default:
		fmt.Printf("no communication\n")
	}
}
