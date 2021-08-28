package main

import "fmt"

func main() {

	Show(1, 2, 3)

	datas := []int{3, 4, 5, 6, 7}
	Show(datas...)

	objects := make([]int, len(datas))
	Show(objects...)
}

func Show(ary ...int) {
	fmt.Println(ary)
}
