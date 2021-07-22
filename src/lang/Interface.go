package main

import "fmt"

type Box interface {
	// 调用1
	call(int) string

	// 调用2
	call1(int, int) string
}

type BoxOne struct {
}

type BoxTwo struct {
}

func (BoxOne) call(int) string {

	fmt.Println("box 1")
	return "box 1"
}

func (BoxTwo) call(int) string {
	fmt.Println("box 2")
	return "box 2"
}

func main() {
	a := new(BoxOne)
	a.call(12)

	b := new(BoxTwo)
	b.call(12)
}
