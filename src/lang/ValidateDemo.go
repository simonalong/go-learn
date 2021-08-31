package main

import (
	"fmt"
	"github.com/SimonAlong/go-learn/src/mikilin"
	"github.com/go-playground/validator/v10"
	"github.com/lunny/log"
)

type ValidateEntity struct {
	Name string `validate:"max=3"`
	Age  int    `validate:"max=3"`
}

type MyEntity struct {
	Name string `matcher:"size=2"`
	Age  int    `matcher:"value={12, 32};range=(12,30]"`
}

func main() {

	validate()
	myTag()
}

func validate() {
	entity := ValidateEntity{Name: "3fsdfasd", Age: 3123123}
	validate := validator.New()
	err := validate.Struct(entity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			return
		}
	}
}

func myTag() {

	myentity := MyEntity{}

	result, err := mikilin.Check(myentity)
	if !result {
		log.Errorf(err.Error())
	}
}
