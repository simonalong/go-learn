package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func main() {
	//str := `value = {12, 3} range=[32, 43]`
	//reg1 := regexp.MustCompile(`(?P<value>[a-zA-Z]+)=(?P<data>\w+@\w+(?:\.\w+)+)`)
	//match := reg1.FindStringSubmatch(str)
	//groupNames := reg1.SubexpNames()
	//fmt.Printf("%v, %v, %d, %d\n", match, groupNames, len(match), len(groupNames))
	//
	//result := make(map[string]string)
	//
	//// 转换为map
	//for i, name := range groupNames {
	//	if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
	//		result[name] = match[i]
	//	}
	//}
	//
	//prettyResult, _ := json.MarshalIndent(result, "", "  ")
	//fmt.Printf("%s\n", prettyResult)

	//regex := regexp.MustCompile(`^.+(\s)*=.+(\s)*$`)
	//sub := regex.FindAllStringSubmatch(str, -1)
	//showJson(sub)

	//
	//
	//var str = "adfasd324232sdfas"
	//
	//var reg = regexp.MustCompile("([a-z]*)(\\d*)([a-z]*)");
	//
	//data := reg.FindAllStringSubmatch(str, -1)
	//fmt.Println(data)

	str := `value = {12, 3};range=[32, 43]`
	var reg = regexp.MustCompile(`^(?<=value)*$`)
	data := reg.FindAllStringSubmatch(str, -1)
	showJson(data)
}

func showJson(object interface{}) {
	bytes, _ := json.Marshal(object)
	fmt.Println(string(bytes))
}
