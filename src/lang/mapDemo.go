package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main1() {
	var a = map[string]string{}
	a["a"] = "a1"
	a["b"] = "b1"
	a["c"] = "c1"

	for _, element := range a {
		fmt.Print(element)
	}

	// a1b1c1
	// c1a1b1
}

func mapShow1() {
	var a = map[string]string{}

	a["a"] = "a1"
	a["b"] = "b1"
	a["c"] = "c1"

	for _, element := range a {
		fmt.Println(element)
	}
}

func main3() {
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

	for k, v := range dataMap {
		fmt.Println(k)
		fmt.Println(v)
	}

	for k := range dataMap {
		fmt.Println(k)
	}

	for _, v := range dataMap {
		fmt.Println(v)
	}

	//dataValue := reflect.ValueOf(dataMap)
	//for _, value := range dataValue.MapKeys() {
	//	//mjson, _ := json.Marshal(value)
	//	//fmt.Println(string(mjson))
	//	fmt.Sprintf("%v", value)
	//}
	//
	//for mapR := dataValue.MapRange(); mapR.Next(); {
	//	key := mapR.Key()
	//	value := mapR.Value()
	//	fmt.Println(key.Interface())
	//	fmt.Println(value.Interface())
	//}
}

func main2() {
	dataMap := map[string]interface{}{}
	//key1, value1 := shortKeyValue("a.b.c", "12")
	//key2, value2 := shortKeyValue("a.b.d.e", "22")
	//key3, value3 := shortKeyValue("a.b.d.f", "33")
	key4, value4 := shortKeyValue("g[0]", "0")
	key5, value5 := shortKeyValue("g[1]", "1")
	key6, value6 := shortKeyValue("g[2]", "2")

	//dataMap = deepPut(dataMap, key1, value1)
	//dataMap = deepPut(dataMap, key2, value2)
	//dataMap = deepPut(dataMap, key3, value3)
	dataMap = deepPut(dataMap, key4, value4)
	dataMap = deepPut(dataMap, key5, value5)
	dataMap = deepPut(dataMap, key6, value6)

	fmt.Println(dataMap)
}

// a.b.c=12转换为，a={b:{c:12}}
func shortKeyValue(key string, value string) (string, interface{}) {
	if strings.Contains(key, ".") {
		innerKeys := strings.SplitN(key, ".", 2)

		newKey, newValue := shortKeyValue(innerKeys[1], value)

		innerValue := map[string]interface{}{}
		innerValue[newKey] = newValue

		return innerKeys[0], innerValue
	} else if strings.Contains(key, "[") && strings.HasSuffix(key, "]") {
		// todo
		return key, value
	} else {
		return key, value
	}
}

func deepPut(dataMap map[string]interface{}, key string, value interface{}) map[string]interface{} {
	mapValue, exist := dataMap[key]
	if !exist {
		if strings.Contains(key, "[") && strings.HasSuffix(key, "]") {

		} else {
			dataMap[key] = value
		}
	} else {
		if reflect.Map == reflect.TypeOf(value).Kind() {
			leftMap := mapValue.(map[string]interface{})
			rightMap := value.(map[string]interface{})

			for rightMapKey := range rightMap {
				leftMap = deepPut(leftMap, rightMapKey, rightMap[rightMapKey])
			}
			dataMap[key] = leftMap
		}
	}

	return dataMap
}

func main() {
	//fmt.Println(peelArray("e"))

	str := "- d: 1\n- d: 2\n- d: 3\n- d: 4"

	fmt.Println(YamlToList(str))


	var dataMap = map[string]*[]int{}
	var inner = []int{}
	inner = append(inner, 1)
	inner = append(inner, 3)
	inner = append(inner, 43)
	inner = append(inner, 2)

	dataMap["a"] = &inner

	value, exist := dataMap["a"]
	if exist {
		*value = append(*value, 554)
	} else {
		dataMap["a"] = &inner
	}

	fmt.Println(dataMap)
}

var rangePattern = regexp.MustCompile("^(.*)\\[(\\d*)\\]$")

func peelArray(nodeName string) (string, int) {
	var index = -1
	var name = nodeName
	var err error

	subData := rangePattern.FindAllStringSubmatch(nodeName, -1)
	if len(subData) > 0 {
		name = subData[0][1]
		indexStr := subData[0][2]
		if "" != indexStr {
			index, err = strconv.Atoi(indexStr)
			if err != nil {
				log.Fatalf("解析错误, nodeName=" + nodeName)
				return "", -1
			}
		}
	}
	return name, index
}

func YamlToList(contentOfYaml string) ([]interface{}, error) {
	resultMap := []interface{}{}
	err := yaml.Unmarshal([]byte(contentOfYaml), &resultMap)
	if err != nil {
		log.Fatalf("YamlToList, error: %v, content: %v", err, contentOfYaml)
		return nil, err
	}

	return resultMap, nil
}
