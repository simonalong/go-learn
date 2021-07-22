package main

import "fmt"

// 数组声明：var variable_name [SIZE] variable_type
// 数组初始化：var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
func main() {

	showArray1()
	arrayInit()
	arrayFor()
	arrayTest2()
}

func showArray1() {
	fmt.Println("======================== showArray ========================")

	// 声明
	var a [2]int
	fmt.Println(a)

	// 设置值
	a[0] = 1
	a[1] = 1
	fmt.Println(a)
}

func arrayInit() {
	fmt.Println("======================== array init ========================")

	// 明确数组
	var a = [2]int{1, 2}
	fmt.Println(a)

	// 不说明类型
	b := [2]int{3, 4}
	fmt.Println(b)

	// 不说明数组长度
	c := [...]int{3, 4, 4, 5, 6}
	fmt.Println(c)

	// 不说明数组长度
	d := []int{3, 4, 4, 5, 6, 43, 543, 63}
	fmt.Println(d)

	// 不说明数组长度
	e := []float64{3, 4, 4, 5, 6, 43, 543, 63, 21, 43, 234}
	fmt.Println(e)
}

func arrayFor() {
	fmt.Println("======================== array init ========================")

	a := []int{1, 32, 2, 3234, 234, 23453, 45, 345, 345, 45}
	for index := range a {
		// 这个是索引
		fmt.Println(index)
	}

	for index, data := range a {
		// 这个是索引和数据
		fmt.Println(index, " = ", data)
	}
}

func arrayTest2() {
	fmt.Println("======================== array test2 ========================")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("data", a[1:2])
	fmt.Println("numbers[1:4] ==", a[1:4])

	println(len(a))
	println(cap(a))

}
