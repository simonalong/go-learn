package main

import (
	"fmt"
)

type MyError struct {
	name string
	age  int
}

func (myError *MyError) Error() string {
	fmt.Println(myError.name)
	return "错误了"
}

func main() {
	result, errorMsg := errFun(1)
	if errorMsg == "" {
		fmt.Println("结果", result)
	} else {
		fmt.Println("异常", errorMsg)
	}
}

func errFun(a int) (string, string) {
	if a < 10 {
		myError := MyError{"小于10了", 12}
		return "error", myError.Error()
	}
	return "", ""
}
