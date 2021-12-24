package main

import (
	"fmt"
	"sort"
)

// 数组声明：var variable_name [SIZE] variable_type
// 数组初始化：var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
func main() {

	showArray1()
	arrayInit()
	arrayFor()
	arrayTest2()
	test()
	testSort()
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

	// data [2]
	fmt.Println("data", a[1:2])
	// numbers[1:4] == [2 3 4]
	fmt.Println("numbers[1:4] ==", a[1:4])

	println(len(a))
	println(cap(a))
}

func arrayTest3() {
	fmt.Println("======================== array test3 ========================")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("data", a[1:2])
	fmt.Println("numbers[1:4] ==", a[1:4])

	println(len(a))
	println(cap(a))
}

func test() {
	fmt.Println("======================== slice ========================")
	//切片初始化，其实也就是不指定具体大小的都可以称之为切片
	//
	//
	//s := []int{1,2,4}
	//fmt.Println(s)
	//
	////新的切片
	//ss:= s[:]
	//fmt.Println(ss)
	//
	////创建切片
	//sss := s[0:1]
	//fmt.Println(sss)
	//
	////创建切片
	//s1 := s[:2]
	//fmt.Println(s1)
	//
	////创建切片
	//s2 := s[1:]
	//fmt.Println(s2)
	//
	////创建切片，通过内置函数make进行初始化切片s，指定的是int类型
	//s3 := make([]int, 2, 12)
	//fmt.Println(s3)
	//fmt.Println(s3)
	//
	////系统内置函数len和cap 这两个函数是用于计算切片的当前长度和容量
	//fmt.Println(len(s3), cap(s3))
	//
	////模拟切片超过容量情况
	//s4 := make([]int, 2,2)
	//fmt.Println(cap(s4))
	//ints := append(s4, 2)
	//fmt.Println(s4)
	//fmt.Println(ints)
	//fmt.Println(cap(s4))
	//
	////切片在未初始化之前是空的是nil
	//var numbers []int
	//if numbers == nil{
	//	fmt.Println("空")
	//}else{
	//	fmt.Println(numbers)
	//}
	//
	//切片的append函数，可以增加多个元素
	var nums []int
	nums = append(nums, 0)
	nums = append(nums, 1)
	nums = append(nums, 1, 2, 3, 2)
	fmt.Println(nums)
	//
	////创建新的切片，但是这个里面是没有数据的
	//nums2 := make([]int, len(nums), cap(nums))
	//fmt.Println(nums2, len(nums2), cap(nums2))
	//
	////切片的copy函数，拷贝内容到nums中
	//copy(nums2, nums)
	//fmt.Println(nums2)

	// 这个算是数组
	ss := []int{}[:]
	ss = append(ss, 1)
	ss = append(ss, 2)
	ss = append(ss, 3)
	fmt.Println(ss)
}

type Node struct {
	name string
	age  int
}

func testSort() {
	fmt.Println("======================== testSort ========================")
	node1 := Node{age: 12}
	node2 := Node{age: 1}
	node3 := Node{age: 123}

	nodes := []Node{}
	nodes = append(nodes, node1)
	nodes = append(nodes, node2)
	nodes = append(nodes, node3)

	// 降序
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].age > nodes[j].age
	})

	fmt.Println(nodes)
}
