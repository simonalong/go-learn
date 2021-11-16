package main

import "fmt"

type ValidateEntity struct {
	Name string
	age  int `validate:"max=3"`
}

type MyEntity struct {
	Name string `matcher:"size=2"`
	Age  int    `matcher:"value={12, 32};range=(12,30]"`
}

func main() {

	validate()
	//myTag()
}

func validate() {
	data := 12

	n := data /2
	fmt.Println(n)
}
