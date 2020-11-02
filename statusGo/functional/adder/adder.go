package adder

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 正统函数式编程写法
type iAdder func(int)(int,iAdder)

func adder2(base int) iAdder{
	return func(v int) (int, iAdder) {
		return base+v,adder2(base+v)
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}


func main() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s,a = a(i)
		fmt.Printf("1+...+%d = %d\n", i, s)
	}
}

// 1.函数式编程
// 函数是一等公民：参数，变量，返回值都可以是函数
// 高阶函数
// 函数 -> 闭包
// 2."正统"函数式编程（不强求）
// 不可变性：不能有状态，只有常量和函数
// 函数只能有一个参数





// 闭包
// 函数体： 局部变量、 自由变量 -> 连接关系




//python 中的闭包， nonlocal 定义分本地变量  原生支持闭包；使用__closure__来查看闭包内容
