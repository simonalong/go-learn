package main

import (
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/compiler"
	"github.com/antonmedv/expr/parser"
)

func main() {
	env := map[string]interface{}{
		"greet":   "Hello, %v",
		"data":    "你好#root.Name还有#current",
		"names":   []string{"world", "you"},
		"sprintf": fmt.Sprintf,
		"println": fmt.Print,
		"root":    Data{Name: "nihao", Age: 21},
		"current": 12,
	}

	//code := `sprintf("你%v好%v还有", root.Name, current)`
	//code := `sprintf("名字%v不合法，还有当前的值%v不满足要求，还有%v is not Empty", root.Age, current, root.Name)`
	code := "名字#root.Age不合法，还有当前的值#current不满足要求，还有#root.Name is not Empty"

	errMsg := errMsgChange(code)

	//fmt.Println(errMsg)

	tree, err := parser.Parse(errMsg)
	if err != nil {
		panic(err)
	}

	program, err := compiler.Compile(tree, nil)
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
	//fmt.Println(program.Disassemble())
}

type Data struct {
	Name string
	Age  int
}

// 数据#root.Age的名字#current不合法
// sprintf("数据%v的名字%v不合法", root.Age, current)
//func chg(str string) string {
//
//}
