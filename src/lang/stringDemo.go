package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "21203129"
	fmt.Println(str)
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Trim(str, " "))
	fmt.Println(strings.Replace(str, "2", "8", 2))
	//
	//fmt.Println(strings.Index(str, "="))
	//index := strings.Index(str, "o")
	//fmt.Println(str[:index])
	//fmt.Println(str[index:])
	//
	//ints := []int{23, 1, 23, 15, 2, 312, 42}
	//sort.Ints(ints)
	//
	//fmt.Println(ints)
	//
	//ints = []int{}
	//fmt.Println(ints)

	//for i, data := range str {
	//	r, _ := strconv.Atoi(string(data))
	//	fmt.Println(i, " ", r)
	//}

	//fmt.Println(str[:len(str)-1])

	//i := interp.New(interp.Options{})
	//
	//i.Use(stdlib.Symbols)
	//
	//i.Use(interp.Exports{
	//	"a": {
	//		"val": reflect.ValueOf(int64(11)),
	//	},
	//})
	//
	//
	//
	//_, err := i.Eval(`import "fmt"`)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = i.Eval(`fmt.Println(a.val)`)
	//if err != nil {
	//	panic(err)
	//}

	test1()
}
func test1() {
	strings.Replace()
}

// 将其中的root.xx和current生成对应的占位符和sprintf字段，比如：数据#root.Age的名字#current不合法，转换为：sprintf("数据%v的名字%v不合法", root.Age, current)
func errMsgToTemplate(errMsg string) string {

	strings.Replace()
}

func Replace(s, old, new string, n int) string {
	if old == new || n == 0 {
		return s // avoid allocation
	}

	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return s // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}

	// Apply replacements to buffer.
	var b strings.Builder
	b.Grow(len(s) + n*(len(new)-len(old)))
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		b.WriteString(s[start:j])
		b.WriteString(new)
		start = j + len(old)
	}
	b.WriteString(s[start:])
	return b.String()
}
