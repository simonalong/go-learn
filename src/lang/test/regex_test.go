package test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
)

func TestRegex1(t *testing.T) {
	str := "dataBaseUserHao"
	reg, err := regexp.Compile("\\B[A-Z]")
	if err != nil {
		return
	}

	subIndex := reg.FindAllStringSubmatchIndex(str, -1)
	var lastIndex = 0
	var result = ""
	for i := 0; i < len(subIndex); i++ {
		result = str[lastIndex:subIndex[i][0]]
		result += str[subIndex[i][0]:subIndex[i][1]]
		lastIndex = subIndex[i][1]
	}
	result += str[lastIndex:]
}

//func TestRegex2(t *testing.T) {
//
//	//fmt.Printf("%s\n", groupNames)
//
//	//result := make(map[string]string)
//	//
//	// 转换为map
//	//for i, name := range groupNames {
//	//	//if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
//	//	//	result[name] = match[i]
//	//	//}
//	//}
//	//
//	//prettyResult, _ := json.MarshalIndent(result, "", "  ")
//	//fmt.Printf("%s\n", prettyResult)
//
//	regex := regexp.MustCompile(`^.+(\s)*=.+(\s)*$`)
//	sub := regex.FindAllStringSubmatch(str, -1)
//	showJson(sub)
//
//	//
//	//
//	//var str = "adfasd324232sdfas"
//	//
//	//var reg = regexp.MustCompile("([a-z]*)(\\d*)([a-z]*)");
//	//
//	//data := reg.FindAllStringSubmatch(str, -1)
//	//fmt.Println(data)
//
//	//str := "zhouOKzhen"
//	//reg,err := regexp.Compile("^\\d+$")
//	//if err != nil {
//	//	fmt.Errorf(err.Error())
//	//	fmt.Println(err.Error())
//	//}
//	//
//	//fmt.Println(reg.MatchString(fmt.Sprintf("%v", 12)))
//	////fmt.Println(reg.String())
//
//	//reg.MatchString()
//	//data := reg.FindAllStringSubmatch(str, -1)
//	//showJson(data)
//
//	//reg := regexp.MustCompile("^(\\(|\\[)(.*),(\\s)*(.*)(\\)|\\])$");
//	//fmt.Println(reg.MatchString("[12, 31]"))
//	//datas := reg.FindAllStringSubmatch("[12, 31]", 2)
//	//for _, data := range datas {
//	//	fmt.Println("====")
//	//	for _, dat := range data {
//	//		fmt.Println(dat)
//	//	}
//	//}
//	//// ====
//	////[12, 31]
//	////[
//	////12
//	////
//	////31
//	////]
//
//	//
//	//reg := regexp.MustCompile("^(\\(|\\[)(.*),(\\s)*(.*)(\\)|\\])$");
//	////fmt.Println(reg.MatchString("[12, 31]"))
//	//datas := reg.FindAllStringSubmatch("fs", 312)
//	////for _, data := range datas {
//	////	for index, dat := range data {
//	////		//fmt.Println(fmt.Sprintf("%v = %v",index, dat))
//	////	}
//	////}
//	//fmt.Println(datas)
//	//fmt.Println(fmt.Sprintf("1 = %v", datas[0][1]))
//	//fmt.Println(fmt.Sprintf("2 = %v", datas[0][2]))
//	//fmt.Println(fmt.Sprintf("3 = %v", datas[0][3]))
//	//fmt.Println(fmt.Sprintf("4 = %v", datas[0][4]))
//	//fmt.Println(fmt.Sprintf("5 = %v", datas[0][5]))
//	//// ====
//	////[12, 31]
//	////[
//	////12
//	////
//	////31
//	////]
//
//	//var rangeRegex = regexp.MustCompile("^(\\(|\\[)(.*)(,|，)(\\s)*(.*)(\\)|\\])$")
//	//subData := rangeRegex.FindAllStringSubmatch("[1，2]", 1)
//	//if len(subData) > 0 {
//	//	//beginAli := subData[0][1]
//	//	//begin := subData[0][2]
//	//	//end := subData[0][4]
//	//	//endAli := subData[0][5]
//	//	for _, data := range subData[0] {
//	//		fmt.Println(data)
//	//	}
//	//}
//	//fmt.Println("daasdf")
//
//	//
//	//var digitRegex = "^(0)|^[-+]?([1-9]+\\d*|0\\.(\\d*)|[1-9]\\d*\\.(\\d*))$"
//	//
//	//// 返回 false
//	//fmt.Println(regexp.MatchString(digitRegex, "2019-07-13 12:00:23.321"))
//	//// 返回true
//	//fmt.Println(regexp.MatchString(digitRegex, "12.321"))
//	//// 返回true
//	//fmt.Println(regexp.MatchString(digitRegex, "0"))
//	//// 返回true
//	//fmt.Println(regexp.MatchString(digitRegex, "-12.98"))
//	//// 返回true
//	//fmt.Println(regexp.MatchString(digitRegex, "+12321"))
//	//
//	//reg := regexp.MustCompile(digitRegex)
//	//datas := reg.FindAllStringSubmatch("2019-07-13 12:00:23.321", -1)[0]
//	//
//	//for _, data := range datas {
//	//	fmt.Println(data)
//	//}
//	var digitRegex = "^([-+])?(\\d*y)?(\\d*M)?(\\d*d)?(\\d*H|\\d*h)?(\\d*m)?(\\d*s)?$"
//
//	fmt.Println(regexp.MatchString(digitRegex, "-12.98"))
//	fmt.Println(regexp.MatchString(digitRegex, "2019-07-13 12:00:23.321"))
//	fmt.Println(regexp.MatchString(digitRegex, ""))
//
//	years, _ := strconv.Atoi(fmt.Sprintf("%v%v", "-", "1"))
//	fmt.Println(years)
//
//	var timePlusRegex = regexp.MustCompile("^([-+])?(\\d*y)?(\\d*M)?(\\d*d)?(\\d*H|\\d*h)?(\\d*m)?(\\d*s)?$")
//	datas := timePlusRegex.FindAllStringSubmatch("-1y2M3d4h5m6s", -1)
//	for i, data := range datas[0] {
//		fmt.Println(fmt.Sprintf("%v = %v", i, data))
//	}
//
//	s, _ := time.ParseDuration("-24h1m")
//	fmt.Println(time.Now().Add(s))
//}

func showJson(object interface{}) {
	bytes, _ := json.Marshal(object)
	fmt.Println(string(bytes))
}
