package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]			//读输入的文件
	if len(files) < 0 {
		countLines(os.Stdin, counts)    //若文件不存在，则读输入命令行数据， Crtl + D 运行结束 return结果
	} else {
		for _, arg := range files {     //进行文件进行数据读取，可多个;arg表示文件名
			f, err := os.Open(arg)  //打开文件操作
			if err != nil {         //常规的非true判定
				fmt.Fprintf(os.Stderr,"dup2:%v\n", err) 
				continue
			}
			countLines(f, counts)  //函数引用
			f.Close()              //文件关闭
		}
	}

//循环读取counts map
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, files[0])
		}
	}
}
//跟进课程需求，改进了输出时添加文件名，按照所学的知识若输入单个文件，则可输出文件名。后续使用接口，就可以对多个文件名进行指定输出

//重复计数
func countLines(f *os.File, counts map[string]int) {  //打开的文件用 *os.File 表示
	input := bufio.NewScanner(f)
	for input.Scan() {		     //重点操作
		counts[input.Text()]++
	}
}
