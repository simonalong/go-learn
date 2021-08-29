package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}

func Test(a, b string) (string, error) {
	return a + " " + b, nil
}

func main() {
	m := map[string]interface{}{"test": Test}
	ret, err := Call(m, "test", "hello", "world")
	if err != nil {
		fmt.Println("method invoke error:", err)
	}
	fmt.Println(ret)
}
