package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ClsDemo struct {
	Name string
	Age  int
}

type ClsArrayDemo struct {
	Name [2][3]string
}

type ClsMapDemo struct {
	Data  map[string]float32
	Data1 map[int]complex64
	Data2 map[string]int
}

type ClsDemoPtr struct {
	Name *string
	Age  *int
}

func (c ClsDemo) Fun11(name string) {
	c.Name = name
}

func (c ClsDemo) Fun21(age int) {
	c.Age = age
}

func (c *ClsDemo) Fun1(name string) {
	c.Name = name
}

func (c *ClsDemo) Fun2(age int) {
	c.Age = age
}

type ValueInnerEntity struct {
	Name string `match:"value={inner_zhou, inner_宋江}"`
	Age  int    `match:"value={2212, 2213}"`
}

type ValueArrayEntity struct {
	Inner [3]ValueInnerEntity `match:"check"`
}

func main() {

	var value ValueArrayEntity
	innerArray := [3]ValueInnerEntity{}

	innerArray[0] = ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	innerArray[1] = ValueInnerEntity{Age: 2213, Name: "inner_zhou"}
	innerArray[2] = ValueInnerEntity{Age: 2214, Name: "inner_宋江"}
	value.Inner = innerArray

	names := [5]string{}
	names[0] = "a"
	names[1] = "b"

	dataValue := reflect.ValueOf(value)
	fieldValue := dataValue.Field(0)
	for arrayIndex := 0; arrayIndex < fieldValue.Len(); arrayIndex++ {
		fieldValueItem := fieldValue.Index(arrayIndex)
		fmt.Println(fieldValueItem)
	}

	//fmt.Println(reflect.ValueOf(value).Kind().String())
	//fmt.Println(reflect.ValueOf(value).Len())
	fmt.Println(reflect.ValueOf(value).Field(0).Index(0))

	//dataType := reflect.TypeOf(data)
	//dataValue := reflect.ValueOf(data)
	//fmt.Printf(reflect.TypeOf(data).Kind().String())

	//field := reflect.TypeOf(data).Field(0)

	//fmt.Println(dataValue.Field(0))
	//fmt.Println(dataValue.Field(0).Kind().String())

	//field.Type.Len()
	//field.Type.In()
	//fmt.Println(dataType.Field(0).Type.Len())
	//fmt.Println(dataType.Field(0).Type)
	//fmt.Println(dataType.Field(0).Type.Kind())
	//fmt.Println(dataType.Field(0).Type.Name())
	//fmt.Println(dataType.Field(0).Type.PkgPath())
	//fmt.Println(dataValue.Field(0).Type().Elem())

	//
	//
	//
	//
	//
	//
	//
	//var datas [3]ClsArrayDemo
	//datas[0] = "a"
	//datas[1] = "b"
	//datas[2] = "c"
	//
	//cls := ClsDemo{Name: "zhou"}
	//
	//fmt.Println(datas)
	//fmt.Println(reflect.TypeOf(datas).Kind().String())
	////dataValue := reflect.ValueOf(datas)
	//fmt.Println(reflect.TypeOf(cls).PkgPath())

	//
	//var dataList []string
	//dataList = append(dataList, "a")
	//dataList = append(dataList, "b")
	//dataList = append(dataList, "c")
	//
	//fmt.Println(dataList)
	//fmt.Println(reflect.TypeOf(dataList).Kind().String())
	//
	//cls := ClsDemo{Name: "zhou"}
	//
	//// main.ClsDemo
	//fmt.Println(reflect.TypeOf(cls).String())
	//// ClsDemo
	//fmt.Println(reflect.TypeOf(cls).Name())
	//
	//objValue := reflect.ValueOf(cls)
	////objType := reflect.TypeOf(cls)
	//fieldValue := objValue.FieldByName("Name")
	//
	//ShowData(fieldValue.NumField())

	//objType.FieldAlign()
	//
	//fmt.Printf("%v", fieldValue.IsNil())
	//fmt.Printf("%v", fieldValue.IsValid())
	//
	//myValue, _ := json.Marshal(fieldValue.Interface())
	//
	//fmt.Printf("%v", string(myValue))
	//fmt.Println("=======")
	////cls.Fun1("chg")
	//
	//// 显示带*的，也显示不带星号的
	////Fun1 func(*main.ClsDemo, string)
	////Fun11 func(*main.ClsDemo, string)
	////Fun2 func(*main.ClsDemo, int)
	////Fun21 func(*main.ClsDemo, int)
	//method(&cls)
	//fmt.Println("=======")
	//// 只显示不带星号的
	////Fun11 func(main.ClsDemo, string)
	////Fun21 func(main.ClsDemo, int)
	//method(cls)
	//
	//fmt.Printf(cls.Name)
	//
	//element()
}

func ShowData(obj ...interface{}) {
	var values []interface{}
	for _, data := range obj {
		myValue, _ := json.Marshal(data)
		values = append(values, string(myValue))
	}
	fmt.Printf(fmt.Sprintf("核查错误：%v", values...))
}

func method(obj interface{}) {
	objType := reflect.TypeOf(obj)

	for i := 0; i < objType.NumMethod(); i++ {
		method := objType.Method(i)
		fmt.Println("index = ", method.Index)
		fmt.Println("Name = ", method.Name)
		fmt.Println("type = ", method.Type)
		fmt.Println("pkgPath = ", method.PkgPath)
		fmt.Println("func = ", method.Func)
		fmt.Println("---------")
	}

	objValue := reflect.ValueOf(obj)
	fmt.Printf("%v", objValue.Interface())
	//
	//data := "sdfsdf"
	//values := make([]reflect.Value, 1)
	//values[0] = reflect.ValueOf(data)
	//valueRun := objValue.MethodByName("Fun11")
	//valueRun.Call(values)
}

func field() {

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

func test3() {
	// ======= 普通占位符 =======
	data1 := ClsDemo{Name: "nihao", Age: 32}
	// {nihao 32}
	fmt.Printf("v = %v\n", data1)
	// {Name:nihao Age:32}
	fmt.Printf("+v = %+v\n", data1)
	//  main.ClsDemo{Name:"nihao", Age:32}
	fmt.Printf("#v = %#v\n", data1)
	// main.ClsDemo
	fmt.Printf("T = %T\n", data1)
	// 字面上的%百分号
	fmt.Printf("%%\n")

	// ======= boolean占位符 =======
	// true
	fmt.Printf("%t\n", true)

	// ======= 整数占位符 =======
	// 二进制：100000
	fmt.Printf("b = %b\n", data1.Age)
	// Unicode码点表示字符
	fmt.Printf("c = %c\n", 0x4E2D)
	// 十进制 32
	fmt.Printf("d = %d\n", data1.Age)
	// 八进制 40
	fmt.Printf("o = %o\n", data1.Age)
	// 单引号围绕的Unicode码
	fmt.Printf("q = %q\n", 0x4E2D)
	// 十六进制，小写a~f，ff
	fmt.Printf("x = %x\n", 255)
	// 十六进制，大写a~f，FF
	fmt.Printf("X = %X\n", 255)
	// Unicode格式，U+0020
	fmt.Printf("U = %U\n", data1.Age)
}

func element() {
	fmt.Println("===== element =====")
	//demo := ClsDemo{}
	//demo.Name = "sdf"
	//
	//demoType := reflect.TypeOf(demo)
	//demoType.Elem()

	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	// name:'' kind:'ptr'
	fmt.Printf("name:'%s' kind:'%s'\n", typeOfCat.Name(), typeOfCat.Kind())
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	// element name: 'cat', element kind: 'struct'
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())

}
