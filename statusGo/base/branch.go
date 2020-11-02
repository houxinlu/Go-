package main

import (
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}


func main() {
	const filename = "abc.txt"
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	fmt.Println(
		grade(1),
		grade(20),
		grade(69),
		grade(85),
		grade(99),
		grade(101),
	)

}

// if 条件里可以赋值
// if 的条件里赋值的变量作用域就在这个if语句里

//switch 后可以没有表达式

//for 
//   sum := 0
//   for i :=1; i <= 100; i++ {
// 	  sum += i
//   }
// for的条件里不需要括号
// for的条件里可以省略初始条件，结束条件，递增表达式
