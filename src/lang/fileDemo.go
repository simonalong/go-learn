package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {

	base := "/home/bob"
	fmt.Println(path.Join(base, "work/go", "src/github.com"))

	fmt.Println(filepath.Abs("../alice"))
}
