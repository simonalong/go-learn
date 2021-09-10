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
		"root":    Data{Name: "nihao"},
		"current": 12,
	}

	code := `sprintf("你%v好%v还有", root.Name, current)`

	tree, err := parser.Parse(code)
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
}

// 数据#root.Age的名字#current不合法
// sprintf("数据%v的名字%v不合法", root.Age, current)
func chg(str string) string {

}
