package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)      //这个很重要，channel是用来在goroutine之间来传参数的
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http") != true {
			url = "https://" + url
		}                  
		go fecth(url, ch)
//		fmt.Println(<-ch)  //放在这里也是可行的
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)   //程序的本质核心数据的传入是从此处开始的
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())  // time.Since(start).Seconds()获取运行start到本位置的运行时间
}

func fecth(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url) //老方法
	if err != nil {
		ch <- fmt.Sprint(err)  //fmt.Sprint() 格式化输出，em... 感觉就是规范么，目前没有深入理解, <- 这个符号类似于传参，很方便
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)  //ioutil.Discard 垃圾桶(类似/dev/null)，但是它能返回repy.Body的字符数
	resp.Body.Close()  //resp.Body 关闭操作（）
	if err != nil {
		ch <- fmt.Sprintf("while reading%s:%v", url, err)  //fmt.Sprinf()和fmt.Sprint()的区别在于表达方式上
		return
	}
	secs := time.Since(start).Seconds() //也可以直接放在 fmt.Sprintf()中，但是这样写美观一些
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
