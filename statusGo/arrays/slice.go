package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("s1=", arr[2:])
	fmt.Println("s2=", arr[:])

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//Slice的扩展
	fmt.Println("Extending slice")
	arr1 := [...]int{0,1,2,3,4,5,6,7}
	s1 = arr1[2:6]  // arr1 [2,3,4,5]
	s2 = s1[3:5]    // arr1 [5,6]
	fmt.Println("s1=",s1)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",s1,len(s1),cap(s1))
	fmt.Println("s2=",s2)
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n",s2,len(s2),cap(s2))

	//append
	s3 := append(s2,10) //实际上，只更改了原 arr 的 7的值，为 10
	s4 := append(s3,11) //超出的值，系统自动放在了别的view
	fmt.Println("s3,s4",s3,s4)
	// s4 no longer view arr1.
	fmt.Println("arr = ",arr1)

}

//slice 不是值传递，它是一个view
//所以实际go语言中，只用slice，基本不用数组

// arr :[...] {0,1,2,3,4,5,6,7}
// s := arr[2:6]
// s[0] = 10
// arr 的值就变成了 0,1,10,3,4,5,6,7

// reslice
// s := arr[2:6]
// s = s[:3]
// s = s[1:]
// s = arr[:]

// Slice 的扩展
// 为什么slice 能扩展，是因为 slice 实现中有三个值， ptr 开头指针，len 长度 ,cap ptr到整个array 的最后的长度。所以 slice可以向后扩展
// 向Slice 添加元素
// 添加元素时，如果超过cap,系统会重新分配更大的底层数组
// 由于值传递关系，必须接受append的返回值
// s = append(s,val)

