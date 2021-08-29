package main

import (
	"fmt"
	"reflect"
	"testing"
)

type clsDemo struct {
	name string
	age  int
}

func main() {
	//a1 := "nihao"
	//showType(a1);
	//

	// ======= 普通占位符 =======
	data1 := clsDemo{name: "nihao", age: 32}
	// {nihao 32}
	fmt.Printf("v = %v\n", data1)
	// {name:nihao age:32}
	fmt.Printf("+v = %+v\n", data1)
	//  main.clsDemo{name:"nihao", age:32}
	fmt.Printf("#v = %#v\n", data1)
	// main.clsDemo
	fmt.Printf("T = %T\n", data1)
	// 字面上的%百分号
	fmt.Printf("%%\n")

	// ======= boolean占位符 =======
	// true
	fmt.Printf("%t\n", true)

	// ======= 整数占位符 =======
	// 二进制：100000
	fmt.Printf("b = %b\n", data1.age)
	// Unicode码点表示字符
	fmt.Printf("c = %c\n", 0x4E2D)
	// 十进制 32
	fmt.Printf("d = %d\n", data1.age)
	// 八进制 40
	fmt.Printf("o = %o\n", data1.age)
	// 单引号围绕的Unicode码
	fmt.Printf("q = %q\n", 0x4E2D)
	// 十六进制，小写a~f，ff
	fmt.Printf("x = %x\n", 255)
	// 十六进制，大写a~f，FF
	fmt.Printf("X = %X\n", 255)
	// Unicode格式，U+0020
	fmt.Printf("U = %U\n", data1.age)

	//%b      无小数部分的，指数为二的幂的科学计数法，
	//        与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
	//%e      科学计数法，例如 -1234.456e+78        Printf("%e", 10.2)     1.020000e+01
	//%E      科学计数法，例如 -1234.456E+78        Printf("%e", 10.2)     1.020000E+01
	//%f      有小数点而无指数，例如 123.456        Printf("%f", 10.2)     10.200000
	//%g      根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出 Printf("%g", 10.20)   10.2
	//%G      根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出 Printf("%G", 10.20+2i) (10.2+2i)

	//show(reflect.TypeOf(data1).Kind())
	//
	//
	//
	////show(reflect.TypeOf(data1).String());
	////show(reflect.ValueOf(data1).String());
	//
	//
	////k1 := reType.Kind()
	////k2 := reVal.Kind()
	////fmt.Printf("k1 = %v type = %T, k2 = %v, type = %T\n", k1, k1, k2, k2)
	//
	//fmt.Println(reflect.ValueOf(data1))

	//data := 1
	//
	//v := reflect.ValueOf(&data)
	//v.Elem().SetInt(100)

	//fmt.Println("%v", data1)

	//var x float64 = 3.14
	//v := reflect.ValueOf(&x)
	//v.Elem().SetFloat(6.28)
	//fmt.Printf("After Set Value is %f", x)

}

func testDemo(t *testing.T) {

}

func show(ary interface{}) {
	t := reflect.TypeOf(ary)
	fmt.Printf("type of a is:%s\n", t)
	//marshal, err := json.Marshal(ary)
	//if err != nil {
	//	return
	//}
	//fmt.Println(fmt.Sprintf("%s", marshal))
	//fmt.Println(fmt.Sprintf("%s", marshal))
}

func showType(ary ...interface{}) {
	for index := range ary {
		show(reflect.TypeOf(ary[index]).Name())
		show(reflect.TypeOf(ary[index]).String())
		show(reflect.TypeOf(ary[index]).Kind().String())
		show(reflect.ValueOf(ary[index]).Kind().String())
		show(reflect.ValueOf(ary[index]).String())
		//show(reflect.TypeOf(ary[index]).Elem().String())
	}
	show("======")
}
