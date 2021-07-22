package main

import "fmt"

func main() {

	mapShow1()
	mapShow2()
}

func mapShow1() {
	fmt.Println("==================== ++ ====================")

	var a = map[string]string{}

	a["a"] = "a1"
	a["b"] = "b1"
	a["c"] = "c1"

	fmt.Println(a)
	fmt.Println(a["a"])
	fmt.Println(a["b"])
}

func mapShow2() {
	fmt.Println("==================== map show ====================")

	var a = make(map[string]string)

	a["a"] = "a1"
	a["b"] = "b1"
	a["c"] = "c1"

	fmt.Println(a)
	fmt.Println(a["a"])
	fmt.Println(a["b"])

	// 循环
	for key := range a {
		fmt.Println(key, a[key])
	}

	fmt.Println("================= 判断是否存在 ==============")
	// 判断是否存在
	value, exist := a["c"]
	fmt.Println(value, " ", exist)

	value2 := a["c"]
	fmt.Println(value2)

	fmt.Println("================= 删除 ==============")
	// 删除数据
	delete(a, "c")
	value, exist = a["c"]
	fmt.Println(value, " ", exist)
	if "" == value {
		fmt.Println(value, " is null")
	}

	value2 = a["c"]
	fmt.Println(value2)

}
