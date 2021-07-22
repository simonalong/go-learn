package main

import "fmt"

func main() {

	showPlus()
}

func showPlus() {
	fmt.Println("==================== ++ ====================")

	a, b := 12, 32

	fmt.Println(a + b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// 不支持--a，以及++a这种
}
