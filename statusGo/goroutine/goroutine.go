package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {	//2.代的是这里的值
			for {
				a[i]++    //3.这里的值跟2一致
				runtime.Gosched()
			}
		}(i) //1.这里的值是传入值
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}


//协程 Coroutine
//轻量级 “线程”
//非抢占式多任务处理，由协程主动交出控制权/但在go1.14版本中，现在可以被异步抢占
//编译器、解释器、虚拟机层面的多任务
//多个协程可能在一个或多个协程上运行
// race condition! -race 数据访问冲突