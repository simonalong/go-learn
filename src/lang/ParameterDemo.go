package main

import "fmt"

func main() {

	show(1, 2, 3)

	datas := []int{3, 4, 5, 6, 7}
	show(datas...)

	objects := make([]int, len(datas))
	show(objects...)
}

func show(ary ...int) {
	fmt.Println(ary)
}
