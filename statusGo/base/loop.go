package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//整数转换二进制数，可以省略初始条件 相当于while
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		result = "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//循环读取文件内容，可省略初始条件和循环条件 相当于while
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}


func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//死循环,for省略所有条件
// func forever() {
// 	for {
// 		fmt.Println("HXL")
// 	}
// }

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(2341141),
		convertToBin(0),
	)
	readFile("abc.txt")
	s := `hxl 123
          xfasd  asd aa "HXL"
			
			qa`
	printFileContents(strings.NewReader(s))
	// forever()
}

// for，if 后面的条件没有括号
// if 条件里也可定义变量
// 没有while
// switch不需要break，也可以直接switch多个条件
// switch{
// case  :
// case  :
// case  :
// }
