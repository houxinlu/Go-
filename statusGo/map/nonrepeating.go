package main

import "fmt"

// 寻找最长不含有重复字符的字串数量
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	} //遍历字符串，将每个字符转换成byte
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabccc"))
	//fmt.Println(lengthOfNonRepeatingSubStr("bb"))
	//fmt.Println(lengthOfNonRepeatingSubStr(""))
	//fmt.Println(lengthOfNonRepeatingSubStr("abcedfg"))
}
