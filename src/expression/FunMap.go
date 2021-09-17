package main

import (
	"fmt"
	"github.com/lunny/log"
	"reflect"
)

var funMap = make(map[string]interface{})

func RegisterFun(funName string, fun interface{}) {
	funValue := reflect.ValueOf(fun)
	if funValue.Kind() != reflect.Func {
		log.Errorf("fun is not fun type")
		return
	}

	if funValue.Type().NumIn() > 2 {
		log.Errorf("the num of argument need to be less than or equal to 2")
		return
	}

	if funValue.Type().NumOut() > 2 {
		log.Errorf("the num of return need to be less than or equal to 2")
		return
	}

	if funValue.Type().NumOut() == 1 {
		if funValue.Type().Out(0).Kind() != reflect.Bool {
			log.Errorf("the type of return must be bool")
			return
		}
	} else {
		if funValue.Type().Out(0).Kind() != reflect.Bool && funValue.Type().Out(1).Kind() != reflect.String {
			log.Errorf("the types of return must be bool and string")
			return
		}
	}

	funMap[funName] = fun
}

func Call(funName string, parameters ...interface{}) (bool, bool, string) {
	if len(parameters) > 2 {
		log.Errorf("the num of parameter must be less than or equal to 2")
		return false, true, ""
	}

	fun, contain := funMap[funName]
	if !contain {
		log.Errorf("the name of fun not find")
		return false, true, ""
	}

	funValue := reflect.ValueOf(fun)

	in := make([]reflect.Value, len(parameters))
	for i, param := range parameters {
		in[i] = reflect.ValueOf(param)
	}

	retValues := funValue.Call(in)
	if len(retValues) == 1 {
		return true, retValues[0].Bool(), ""
	} else {
		return true, retValues[0].Bool(), retValues[1].String()
	}
}

func main() {
	RegisterFun("test", Test)
	validate, match, errMsg := Call("test", "陈真")
	if !validate {
		log.Errorf("check error")
		return
	}

	if !match {
		log.Errorf("match error: %v", errMsg)
	}

	fmt.Println(match)
}

func Test(name string) (bool, string) {
	if name == "zhou" || name == "宋江" {
		return true, ""
	} else {
		return false, "当前的值" + name + "不在合法的值[zhou, 宋江]里面"
	}
}
