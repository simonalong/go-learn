package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {

	mapShow1()
	mapShow2()
	mapShow3()
	mapPertty()
	mapShow4()
}

func mapShow1() {
	fmt.Println("==================== ++ ====================")

	// 定义map结构：key:string, value: string
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

func mapShow3() {
	var dataMap = make(map[string]map[string][]int)

	var innerMap1 = make(map[string][]int)
	innerMap1["inner_a"] = []int{2, 3, 4, 5, 435}
	innerMap1["inner_b"] = []int{2, 3, 4, 5, 435}
	innerMap1["inner_c"] = []int{2, 3, 4, 5, 435}
	dataMap["inner1"] = innerMap1

	var innerMap2 = make(map[string][]int)
	innerMap2["inner_a"] = []int{2, 3, 4, 5, 435}
	innerMap2["inner_b"] = []int{2, 3, 4, 5, 435}
	innerMap2["inner_c"] = []int{2, 3, 4, 5, 435}
	dataMap["inner2"] = innerMap2

	mjson, _ := json.Marshal(dataMap)
	fmt.Println(string(mjson))

	mjson2, _ := json.MarshalIndent(dataMap, "", "\t")
	fmt.Println(string(mjson2))
}

func mapPertty() {
	var dataMap = make(map[string]string)
	dataMap["nihao"] = "hello"
	dataMap["nihaoguodoush"] = "word"
	dataMap["shuia"] = "ok"
	dataMap["zhong"] = "haode"

	// json美化
	bytes, _ := json.MarshalIndent(dataMap, "", "\t")
	fmt.Println(string(bytes))
}

func mapShow4() {
	fmt.Println("================= mapShow4 ==============")
	var dataMap = make(map[string]int)
	dataMap["a"] = 1
	dataMap["c"] = 3

	dataValue := reflect.ValueOf(dataMap)
	for _, value := range dataValue.MapKeys() {
		//mjson, _ := json.Marshal(value)
		//fmt.Println(string(mjson))
		fmt.Sprintf("%v", value)
	}

	for mapR := dataValue.MapRange(); mapR.Next(); {
		key := mapR.Key()
		value := mapR.Value()
		fmt.Println(key.Interface())
		fmt.Println(value.Interface())
	}
}
