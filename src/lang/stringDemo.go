package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "caoda   ok  zhou zhen = yong    "
	//fmt.Println(str)
	//fmt.Println(strings.TrimSpace(str))
	//fmt.Println(strings.Trim(str, " "))
	//fmt.Println(strings.Replace(str, " ", "", -1))
	//
	//fmt.Println(strings.Index(str, "="))
	index := strings.Index(str, "o")
	fmt.Println(str[:index])
	fmt.Println(str[index:])

	ints := []int{23, 1, 23, 15, 2, 312, 42}
	sort.Ints(ints)

	fmt.Println(ints)

	ints = []int{}
	fmt.Println(ints)

}
