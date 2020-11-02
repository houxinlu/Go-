package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "*":
		return a * b, nil
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0,
			fmt.Errorf("unsupported operation: %s", op) //注意这里是 fmt.Errorf()
	}
}

// 13 / 3 = 4...1
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数返回对多个值

//函数式编程，函数作为参数。替换eval  * 难点
func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling %s with %d,%d\n", runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(), a, b)
	return op(a, b)
}

// 可变参数列表,可变参数是一个列表，用for取值
func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main() {
	if result, err := eval(3, 4, "X"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3) //收返回值
	fmt.Println(q, r)

	fmt.Println(apply(func(a,b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))    //这个其实相当于函数头去掉了内置函数，内置函数换成了return func()


	fmt.Println(sumArgs(1, 2, 3, 4, 5))
}


//函数语法要点回顾
//返回值类型写在最后面
//可返回多个值
//函数作为参数
//没有默认参数，可选参数
