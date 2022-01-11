package test

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	fmt.Println(time.Second)
	fmt.Println(time.Second.Nanoseconds())
	fmt.Println(5000000000)
	fmt.Println(5000000000 / time.Second)
}
