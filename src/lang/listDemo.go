package main

import (
	"container/list"
	"fmt"
)

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

	queue := list.New()

	queue.PushBack(1)
	queue.PushBack(3)
	queue.PushBack(6)

	fmt.Println(queue)
	if queue.Len() > 1 {
		queue.Remove(queue.Front())
	}

	fmt.Println(queue.Front().Value)
}
