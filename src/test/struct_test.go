package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Demo1 struct {
	InnerName string
	InnerAge  int
}

type DemoImpl struct {
	Demo1
	Name string
	Age  int
}

func TestDemo1(t *testing.T) {
	demo1 := DemoImpl{Demo1: Demo1{InnerName: "zhou", InnerAge: 12}, Name: "out", Age: 90}
	v := reflect.ValueOf(demo1)
	fmt.Println(v.NumField())

	for index, num := 0, v.NumField(); index < num; index++ {
		fmt.Println(v.Field(index).Interface())
	}
}
