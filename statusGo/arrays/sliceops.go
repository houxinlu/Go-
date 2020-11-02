package main

import (
	"fmt"
)

func printSlices(s []int)  {
	fmt.Printf("%v,len=%d,cap=%d\n",s,len(s),cap(s))

}

func main() {
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{1,2,3,4}
	s2 := make([]int,16)
	s3 := make([]int,10,32)
	printSlices(s2)
	printSlices(s3)

	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlices(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3],s2[4:]...)   //这里特别注意，append()
	printSlices(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("front = ",front)
    printSlices(s2)
	
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail = ",tail)
    printSlices(s2)
}

// slice 当你不断append的时候 会进行扩充cap, 分别是 0，2，4，8，16，32 .. 2的次方
// 建立一个长度限定的slice  s := make([]int,16)   s1 := make([]int,10,32)
