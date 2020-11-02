package main

import (
	"bufio"
	"fmt"
	"os"
	"statusGo/functional/adder"
)

func typeDefer() {
	defer fmt.Println("1")
	defer fmt.Println(2)
	fmt.Println(3)
	return
	defer fmt.Println(4)
}

func writeFile(filename string) {
	//file, err := os.Create(filename) //打开 创建文件
	file,err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE,0666)
	if err != nil {
		//panic(err)
		//fmt.Println("file already exists",err)
		if pathError,ok := err.(*os.PathError); !ok{
			panic(err)
		}else {
			fmt.Printf("%s,%s,%s",pathError.Op,pathError.Path,pathError.Err)
		}
		return
	}
	defer file.Close() //当有需要处理时，及时进行处理就行

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := adder.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}

// fmt.Fprintln(文件写入)  fmt.Fprint() fmt.Fprintln 按照默认格式写入，fmt.Fprintf() 按照format格式说明符将内如格式化写入文件

//defer 先进后出 栈序，defer 不会由于return 和 panic 造成影响
//参数在defer语句时计算
//确保调用在函数结束时发生

//调用defer的时机
//Open/Close
//Lock/Unlock
//PrintHeader/PrintFooter