package main

import (
	"fmt"
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
}
