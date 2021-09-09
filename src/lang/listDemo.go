package main

import "fmt"

func main() {
	//datas := []int{12, 2,4,5}
	//
	//for i, data := range datas {
	//	fmt.Println(data)
	//	fmt.Println(i)
	//}

	myArray := [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	//打印一维数组长度
	data := len(myArray)
	fmt.Println(data)
}
