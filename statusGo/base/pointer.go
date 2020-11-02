package main

import (
	"fmt"
)

//验证指针,使用指针类型进行swap操作，不然swap是其实是不生效的
func swap(a, b *int) {
	*a, *b = *b, *a
}

//这种方式是不用指针的,当然正常写法，这种是最好的
func swap2(n, m int) (int, int) {
	return m, n
}

func main() {
	fmt.Println("Hello word")

	//修改指针值，从2改成3
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	a, b := 3, 4
	m, n := 3, 4
	swap(&a, &b)
	n, m = swap2(m, n)
	fmt.Println(a, b)
	fmt.Println(n, m)
}

//指针不能运算
//Go语音只有值传递一种方式，不过可以通过指针实现引用传递的功能
//参数传递，值传递？引用传递？
//值传递，会重新拷贝，不会改变变量本身的值  a
//引用传递，不会拷贝，会改变变量本身的值   &a
//有个参数传递的内容，可以反复观看  2-7节
