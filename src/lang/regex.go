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

	//fmt.Printf("%s\n", groupNames)

	//result := make(map[string]string)
	//
	// 转换为map
	//for i, name := range groupNames {
	//	//if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
	//	//	result[name] = match[i]
	//	//}
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

	//str := "zhouOKzhen"
	//reg,err := regexp.Compile("^\\d+$")
	//if err != nil {
	//	fmt.Errorf(err.Error())
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(reg.MatchString(fmt.Sprintf("%v", 12)))
	////fmt.Println(reg.String())

	//reg.MatchString()
	//data := reg.FindAllStringSubmatch(str, -1)
	//showJson(data)

	//reg := regexp.MustCompile("^(\\(|\\[)(.*),(\\s)*(.*)(\\)|\\])$");
	//fmt.Println(reg.MatchString("[12, 31]"))
	//datas := reg.FindAllStringSubmatch("[12, 31]", 2)
	//for _, data := range datas {
	//	fmt.Println("====")
	//	for _, dat := range data {
	//		fmt.Println(dat)
	//	}
	//}
	//// ====
	////[12, 31]
	////[
	////12
	////
	////31
	////]

	//
	//reg := regexp.MustCompile("^(\\(|\\[)(.*),(\\s)*(.*)(\\)|\\])$");
	////fmt.Println(reg.MatchString("[12, 31]"))
	//datas := reg.FindAllStringSubmatch("fs", 312)
	////for _, data := range datas {
	////	for index, dat := range data {
	////		//fmt.Println(fmt.Sprintf("%v = %v",index, dat))
	////	}
	////}
	//fmt.Println(datas)
	//fmt.Println(fmt.Sprintf("1 = %v", datas[0][1]))
	//fmt.Println(fmt.Sprintf("2 = %v", datas[0][2]))
	//fmt.Println(fmt.Sprintf("3 = %v", datas[0][3]))
	//fmt.Println(fmt.Sprintf("4 = %v", datas[0][4]))
	//fmt.Println(fmt.Sprintf("5 = %v", datas[0][5]))
	//// ====
	////[12, 31]
	////[
	////12
	////
	////31
	////]

	var rangeRegex = regexp.MustCompile("^(\\(|\\[)(.*)(,|，)(\\s)*(.*)(\\)|\\])$")
	subData := rangeRegex.FindAllStringSubmatch("[1，2]", 1)
	if len(subData) > 0 {
		//beginAli := subData[0][1]
		//begin := subData[0][2]
		//end := subData[0][4]
		//endAli := subData[0][5]
		for _, data := range subData[0] {
			fmt.Println(data)
		}
	}
	fmt.Println("daasdf")

	regexp.MustCompile("^([-+])?(\\d*y)?(\\d*M)?(\\d*d)?(\\d*H|\\d*h)?(\\d*m)?(\\d*s)?$")
}

func showJson(object interface{}) {
	bytes, _ := json.Marshal(object)
	fmt.Println(string(bytes))
}
