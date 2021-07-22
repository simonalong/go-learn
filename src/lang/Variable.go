package main

import (
	"fmt"
	"unsafe"
)

// 变量使用方式：
// 一个变量的声明：var identifier type
// 多个变量的声明：var identifier1, identifier2 type
// 变量声明的初始化：
//     var v_name v_type
//     v_name = value
// 也可以这样：var v_name = value
// 声明新的变量：var_name := value
// 多变量声明和初始化：
//      var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
//      vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

// 这种因式分解关键字的写法一般用于声明全局变量
//var (
//    vname1 v_type1
//    vname2 v_type2
//)

var GTest int

// 常量：const c_name1, c_name2 = value1, value2
func main() {
	GTest = 12
	showVar1()
	showVar2()
	showBool()
	showFloat()
	showComplex()
	showInt()
	showInt2()
	showConst()
	showConstIota()
	showPoint()

	println(GTest)
}

func showVar1() {
	fmt.Println("========= := =========")

	// 变量
	test := 12
	println(test)

	// 多个赋值，这种用于首次初始化，自动识别类型
	test1, test2 := 12, 32
	println(test1, test2)
}

func showVar2() {
	fmt.Println("========= = =========")

	// 变量
	var test1 = 12
	println(test1)
}

func showFloat() {
	fmt.Println("========= float =========")

	// +0.000000e+000
	var c float32
	println(c)

	// +0.000000e+000
	var d float64
	println(d)

	// 4
	fmt.Println(unsafe.Sizeof(c))
	// 8
	fmt.Println(unsafe.Sizeof(d))
}

func showComplex() {
	fmt.Println("========= complex =========")

	var i1 complex64 = 1
	var i2 complex128 = 2

	// 8
	fmt.Println(unsafe.Sizeof(i1))
	// 16
	fmt.Println(unsafe.Sizeof(i2))
}

func showBool() {

	fmt.Println("========= bool =========")
	// 默认false
	var b bool
	println(b)

	// 1
	c := true
	fmt.Println(unsafe.Sizeof(c))
}

// 其中int默认是int64
func showInt() {
	fmt.Println("========= int =========")

	// 默认0
	var a int
	println(a)

	// 默认0
	var e int8
	println(e)

	// 默认0
	var f int16
	println(f)

	// 默认0
	var g int32
	println(g)

	// 默认0
	var h int64
	println(h)

	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	// rune 同int32
	var run2 rune = 4
	var i5 int64 = 5

	// 8
	fmt.Println(unsafe.Sizeof(i1))
	// 1
	fmt.Println(unsafe.Sizeof(i2))
	// 2
	fmt.Println(unsafe.Sizeof(i3))
	// 4
	fmt.Println(unsafe.Sizeof(i4))
	// 4
	fmt.Println(unsafe.Sizeof(run2))
	// 8
	fmt.Println(unsafe.Sizeof(i5))
}

func showInt2() {
	fmt.Println("========= uint =========")

	// 同uint64
	var i1 uint = 1
	var i2 uint8 = 2
	// byte同uint8
	var bt byte = 2
	var i3 uint16 = 3
	var i4 uint32 = 4
	var i5 uint64 = 5
	// 无符号的整数，用于存放指针
	var ptr uintptr = 5

	// 8
	fmt.Println(unsafe.Sizeof(i1))
	// 1
	fmt.Println(unsafe.Sizeof(i2))
	// 1
	fmt.Println(unsafe.Sizeof(bt))
	// 2
	fmt.Println(unsafe.Sizeof(i3))
	// 4
	fmt.Println(unsafe.Sizeof(i4))
	// 8
	fmt.Println(unsafe.Sizeof(i5))
	fmt.Println(unsafe.Sizeof(ptr))
}

// 常量的用法
func showConst() {
	fmt.Println("========= const =========")
	// 不可变数据
	const data int = 12
	println(data)

	//a = 32;
	//print(a)

	// 可以做为枚举，其中分号只是用于换行
	const (
		LOCK   = 0
		UNLOCK = 1
	)

	println(LOCK, UNLOCK)

	// 常量也可以做计算
	const (
		ONE   = "1"
		TWO   = 2
		THREE = 12 + 32
		FOUR  = unsafe.Sizeof(12)
	)
	println(ONE, TWO, THREE, FOUR)
}

// 其中iota这个是一个特殊的常量，在const出现会变成0
func showConstIota() {
	fmt.Println("========= const iota =========")

	// iota这个是归零设置，算是从0自增的一个整数数据，这样跟java中的枚举一样了
	const (
		A = iota
		B = iota
		C = iota
		D = iota
		E = iota
	)
	// 也可以简写，后面的可以不用写
	//const (
	//	A = iota
	//	B
	//	C
	//	D
	//	E
	//)
	// 0 1 2 3 4
	println(A, B, C, D, E)

	// 如果首次设置，则后面的都一样了
	const (
		a1 = 10
		b1
		c1 = iota
		d1
		e1
	)
	// 10 10 2 3 4
	println(a1, b1, c1, d1, e1)

	// 如果首次设置，则后面的都一样了
	const (
		a = 12
		b
		c
		d
		e
	)
	// 12 12 12 12 12
	println(a, b, c, d, e)

	// 移位设置
	const (
		a2 = 1 << iota // 1
		b2             // 2
		c2             // 4
		d2             // 8
		e2             // 16
	)
	// 12 12 12 12 12
	println(a2, b2, c2, d2, e2)
}

// 展示指针
func showPoint() {
	fmt.Println("========= point指针 =========")

	a := 12
	// 声明指针
	var b *int

	// 获取指针对应的值
	b = &a

	// 修改指针对应的值
	*b = 33

	println(a)
}
