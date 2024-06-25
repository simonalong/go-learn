package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "UserName"
	fmt.Println(strings.ToLower(str[:1]) + str[1:])

	datas := make([]int, 2)
	datas = append(datas, 12)
}
