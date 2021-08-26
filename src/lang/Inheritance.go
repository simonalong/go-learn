package main

import "fmt"

type Handler interface {
	handle(string) string
}

type Parent struct {
	name string
	age  int8
}

type Child struct {
	// 这里就是继承父类
	Parent

	name string
}

// 继承接口Handler
func (c *Parent) handle(string) string {
	return c.name
}

// 继承接口Handler，与Parent类相同，则会覆盖父类的接口
func (c *Child) handle(string) string {
	return c.name
}

func main() {
	child := Child{name: "child", Parent: Parent{name: "parent", age: 123}}

	// print child
	fmt.Println(child.handle("haha"))
	// print parent
	fmt.Println(child.Parent.handle("haha"))
}
