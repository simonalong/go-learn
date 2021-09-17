package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

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
	entity := ValidateEntity{Name: "3fsdfasd", age: 3123123}
	validate := validator.New()
	err := validate.Struct(entity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			return
		}
	}
}
