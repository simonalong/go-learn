package main

func main() {

	a1, b1 := fun1()
	a2, b2 := fun2()

	println(a1, b1)
	println(a2, b2)

	entity1 := TestEntity{"change", 12}
	fun3(entity1)

	println(entity1.name)

	fun4(func(a int) string {
		println(a)
		return "ok"
	})
}

// 声明返回值的值
func fun1() (a int, b float32) {
	a = 21
	b = 33.0
	return
}

// 不声明，按照返回
func fun2() (int, float32) {
	return 12, 32.9
}

func fun3(entity1 TestEntity) {

	entity1.name = "change"
	entity1.age = 12
}

type TestEntity struct {
	name string
	age  int
}

func toString(entity TestEntity) string {
	// tdo
	return ""
}

type CallBack func(int) string

func fun4(callback CallBack) {
	result := callback(1)
	println(result)
}
