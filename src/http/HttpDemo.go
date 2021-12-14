package main

import (
	"fmt"
	"log"
)

func main() {
	url := "http://10.30.30.78:29013/api/core/license/info"

	data, err := GetSimpleOfStandard(url)
	if err != nil {
		log.Fatalln("err: %v ", err.Error())
	}
	fmt.Println(string(data))
}
