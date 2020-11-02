package main

import (
	"fmt"
	"math"
)

//在函数外部不能进行 := 进行，但是不是全局变量
var aa = 3
var ss = "kkk"

// 整体进行定义，但是不是全局变量
var (
	xc  = 3
	asd = "xxa"
	nn  = true
)

//定义变量 有初始值
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "dfs"
	b = 5
	fmt.Println(a, b, c, s)
}

//常量定义
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	const (
		A = iota
		B
		_
		D
		C
	)
	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota) //每次左移一位
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(A, B, C, D)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(nn)
	consts()
	enums()
}

//变量定义
// 使用var 关键字， 可放在函数内，或者直接放在包内（但是不是全局变量）
// 使用var() 集中定义
// 使用:=   推荐，只能在函数内使用

//内建变量类型
//bool,string
//(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr 指针
//byte(8位),rune(字符型,32位)
//float32,float64,complex64,complex128(复数)

//常量定义
//普通常量相当于替换
//自增型枚举类型

//变量类型要点回顾
//变量类型写在变量名之后
//编译器可推测变量类型
//没有char,只有rune
//原生支持复数类型
