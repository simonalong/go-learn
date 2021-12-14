package main

import "fmt"

func main() {

	var data = "12"
	// 12
	fmt.Println(fmt.Sprintf("%v", data))

	var entity TestEntityK
	entity.Name = "zhou"
	entity.Age = 12
	// {zhou 12}
	fmt.Println(fmt.Sprintf("%v", entity))
	// {Name:zhou Age:12}
	fmt.Println(fmt.Sprintf("%+v", entity))
	// main.TestEntityK{Name:"zhou", Age:12}
	fmt.Println(fmt.Sprintf("%#v", entity))
	// main.TestEntityK
	fmt.Println(fmt.Sprintf("%T", entity))

	// 输出%自己
	fmt.Println(fmt.Sprintf("%%"))

	// {%!b(string=zhou) 1100}
	fmt.Println(fmt.Sprintf("%b", entity))

	// {%!o(string=zhou) 14}
	fmt.Println(fmt.Sprintf("%o", entity))

	// {%!d(string=zhou) 12}
	fmt.Println(fmt.Sprintf("%d", entity))

	// {7a686f75 c}
	fmt.Println(fmt.Sprintf("%x", entity))

	// {7A686F75 C}
	fmt.Println(fmt.Sprintf("%X", entity))

	// {%!U(string=zhou) U+000C}
	fmt.Println(fmt.Sprintf("%U", entity))

	// {%!f(string=zhou) %!f(int=12)}
	fmt.Println(fmt.Sprintf("%f", entity))

	// 显示的指针的值
	fmt.Println(fmt.Sprintf("%p", &entity))
}

type TestEntityK struct {
	Name string
	Age  int
}
