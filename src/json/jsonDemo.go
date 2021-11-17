package main

import (
	"encoding/json"
	"fmt"
)

type Entity struct {
	Name string
	Age  int8
}

func main() {
	entity := Entity{"nihao", 12}
	fmt.Println(json.Marshal(entity))
}
