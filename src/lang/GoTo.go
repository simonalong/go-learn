package main

import "fmt"

func main() {

	lableFun1(12)
}

func lableFun1(a int) {
LABLE1:
	if a > 100 {
		fmt.Println("数据变大")
	} else {
		fmt.Println("数据正常")
		a = a + 1
		goto LABLE1
	}
}
