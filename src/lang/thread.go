package main

import "fmt"

// go 特殊标识符直接开启协程
func main() {
	dataList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 不带缓冲区的数据
	ch := make(chan int)
	go sum(dataList[:len(dataList)/2], ch)
	go sum(dataList[len(dataList)/2:], ch)

	var left = <-ch
	var right = <-ch
	fmt.Println(left, right, left+right)

	sum2()
	show3()
}

func sum(dataList []int, ch chan int) {

	sum := 0
	for _, data := range dataList {
		sum += data
	}
	ch <- sum
}

func sum2() {
	// 带缓冲区的数据
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch, " ", <-ch)
}

func show3() {
	fmt.Println("================= show3 ==============")
	ch := make(chan int, 100)
	go threadFun1(ch)

	for index := range ch {
		fmt.Println(index)
	}
}

func threadFun1(ch chan int) {
	dataList := []int{10, 25, 43, 234, 45}

	for _, data := range dataList {
		ch <- data
	}

	// 该行代码必须有，否则就会出现错误：fatal error: all goroutines are asleep - deadlock!
	close(ch)
}
