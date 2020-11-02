package main

import (
	"fmt"
)

// 遍历,调用 func 会拷贝,不实用
func printArray(arr [5]int) {
	arr[0] = 121
	for i, v := range arr {
		fmt.Println(i,v)
	}
}
// 不实用
func printArray2(arr *[5]int) {
	arr[0] = 121
	for i, v := range arr {
		fmt.Println(i,v)
	}
}

// 通常使用这种方式，slice
func printSlice(arr []int) {
	arr[1] = 212
	for i, v := range arr {
		fmt.Println(i,v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 4}
	arr3 := [...]int{2, 3, 4, 6, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	fmt.Println("arr1 Array")
	fmt.Println(arr1)
	printArray(arr3)
	//int类型不通，不能传递
	//printArray(arr2)

	printArray2(&arr1)
	fmt.Println("arr1 Array2")
	fmt.Println(arr1)

	//printSlice()
	printSlice(arr1[:])
	fmt.Println("arr1 Slice")
	fmt.Println(arr1)

}

//数组是值类型
//[10]int 和 [20]int 是不同类型
//调用func f(arr [10]int) 会拷贝数组
// 例如，在printArray()中改变某个位的值，只能在printArray()函数中生效，main中被数组的位，是不会被改变的
// 当然，我们可以用指针实现 在函数中，进行值的改变 
//在go语言中一般不直接使用数组,使用slice()
