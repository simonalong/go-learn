package main

import "fmt"

func main() {

	showFor1()
	showFor2()
}

func showFor1() {
	println("================ for1 ================")

	// for init; condition; post { }
	// for condition { }
	// for { }

	// 普通的传递
	count := 12
	for index := 0; index <= 10; index++ {
		count++
	}
	println(count)

	// 索引的向后传递
	index := 0
	for ; index <= 10; index++ {
		count++
	}
	println(count)

	// 与while差不多
	for index <= 20 {
		index++
		count++
	}
	println(count)
}

// 数组的循环
func showFor2() {
	println("================ for2 ================")

	a := []int{12, 32}

	for index, data := range a {
		fmt.Println(index, " = ", data)
	}
}
