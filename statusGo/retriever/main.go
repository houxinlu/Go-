package main

import (
	"fmt"
	"statusGo/retriever/mock"
	real2 "statusGo/retriever/real"
	"time"
)


type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("http://www.imocc.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imocc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting",r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"This is a fake Url"}
	//r = mock.Retriever{"This is a fake Url"}
	r = &retriever
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)

	//Type assertion  接口变量内在含义
	//real2Retriever := r.(*real2.Retriever)
	//fmt.Println(real2Retriever.Timeout)
	if mockretriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockretriever.Contents)
	} else {
		fmt.Println("not mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	//fmt.Println(download(r))
}

// 接口变量里面有什么
// 接口变量[实现者的类型  实现者的指针]
// 接口变量 指向  实现着
// 接口变量自带指针
// 接口变量同样采用值传递，几乎不需要使用接口的指针
// 指针接收者实现只能以指针方式使用，值接收者都可

// 查看接口变量
// 表示任何类型： interface{}
// Type Assertion
// Type Switch

// interface{} 可以指代任何类型

// 接口的组合，interface{} 可以组合接口的功能给实现的时候使用
// 常用的系统接口
// Stringer()  Reader/Write
