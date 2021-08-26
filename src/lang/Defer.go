package main

func main() {

	// 最后执行
	defer end1()
	// 后执行
	defer end2()
	// 先执行
	defer end3()
	println("hello")
}

func end1() {
	println("end1")
}

func end2() {
	println("end2")
}

func end3() {
	println("end3")
}
