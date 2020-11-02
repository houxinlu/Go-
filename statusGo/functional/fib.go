package main

import (
	"bufio"
	"fmt"
	"io"
	"statusGo/functional/adder"
	"strings"
)

//1,1,2,3,5,8,13
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}



func main() {
	var f intGen
	f = adder.Fibonacci()
	printFileContents(f)
	//fmt.Println(f()) //1
	//fmt.Println(f()) //1
	//fmt.Println(f()) //2
	//fmt.Println(f()) //3
	//fmt.Println(f()) //5
	//fmt.Println(f()) //8
	//fmt.Println(f()) //13
	//fmt.Println(f()) //12
}

// 例1：斐波那契数列
// 例2：为函数实现接口
// 例3：使用函数遍历二叉树
// go语言闭包的应用
// 更为自然，不需要修饰如何访问自由变量
// 没有Lambda表达式，但是有匿名函数
//
