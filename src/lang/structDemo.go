package main

import (
	"fmt"
	"time"
)

type demo struct {
	name string `form:"Name" binding:"required"`
}

type Demo1 struct {
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

	var demo = Demo1{}
	var demo1 = Demo1{}
	if demo == demo1 {
		fmt.Println("true")
	}

	d := Demo1{name: "ok"}
	ttt(&d)
	fmt.Println(d)

	for {
		fmt.Println("xxxx")
		time.Sleep(1 * time.Second)
	}
}

func ttt(dd *Demo1) {
	dd.name = "change"
}
