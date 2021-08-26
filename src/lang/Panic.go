package main

import "fmt"

func main() {

	defer func() {
		// recover() 函数获取panic函数抛出的异常
		if totalErr := recover(); totalErr != nil {
			fmt.Println(totalErr)
			return
		}

		fmt.Println("end")
	}()

	// 抛出panic异常
	panicRun()

	fmt.Println("执行")
}

func panicRun() {
	panic("抛出异常了哈")
}
