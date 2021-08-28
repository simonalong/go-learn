package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	show(reflect.Bool.String())

	show(os.Environ())
}

func show(ary ...interface{}) {
	marshal, err := json.Marshal(ary)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("%s", marshal))
}
