package main

import (
	"encoding/json"
	"fmt"
)

// 不添加转换
type Entity struct {
	Name string
	Age  int8
}

// 添加字段的转换
type Entity2 struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

func main() {
	entity := Entity{"nihao", 12}
	data, _ := json.Marshal(entity)
	// {"Name":"nihao","Age":12}
	fmt.Println(string(data))

	entity2 := Entity2{"nihao", 12}
	data2, _ := json.Marshal(entity2)
	// {"name":"nihao","age":12}
	fmt.Println(string(data2))

	str := "[\"adsfasdf\", \"fff\"]"
	var datas []string
	_ = json.Unmarshal([]byte(str), &datas)
	// {"name":"nihao","age":12}
	fmt.Println(datas)
}
