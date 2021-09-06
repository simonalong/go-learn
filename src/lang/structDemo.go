package main

import (
	"fmt"
	"regexp"
)

type demo struct {
	name string `form:"Name" binding:"required"`
}

type A interface {
	test()
}

type B struct {
	Name string
}

type C struct {
	Age int
}

func (receiver *B) test() {
	fmt.Println("b的test")
}

func (receiver *C) test() {
	fmt.Println("c的test")
}

func main() {
	//datas := []A{}
	//
	//b := B{Name: "zhou"}
	//c := C{Age: 12}
	//datas = append(datas, &b)
	//datas = append(datas, &c)
	//
	//for _, data := range datas {
	//	data.test()
	//}
	//
	//
	//t := time.Now()
	//fmt.Println(reflect.TypeOf(t).String())
	//

	var digitRegex = "^(0)|^[-+]?([1-9]+\\d*|0\\.(\\d*)|[1-9]\\d*\\.(\\d*))$"

	// 返回 false
	fmt.Println(regexp.MatchString(digitRegex, "2019-07-13 12:00:23.321"))
	// 返回true
	fmt.Println(regexp.MatchString(digitRegex, "12.321"))
	// 返回true
	fmt.Println(regexp.MatchString(digitRegex, "0"))
	// 返回true
	fmt.Println(regexp.MatchString(digitRegex, "-12.98"))
	// 返回true
	fmt.Println(regexp.MatchString(digitRegex, "+12321"))

	reg := regexp.MustCompile(digitRegex)
	datas := reg.FindAllStringSubmatch("2019-07-13 12:00:23.321", -1)[0]

	for _, data := range datas {
		fmt.Println(data)
	}
}
