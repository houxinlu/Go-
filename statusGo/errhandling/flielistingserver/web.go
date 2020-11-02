package main

import (
	"log"
	"net/http"
	"os"
	"statusGo/errhandling/flielistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error // 接口，只对error进行处理

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {		//这个地方其实就只对请求进行了出错处理，不附加任何其他逻辑，writer和request没有任何更改
	return func(writer http.ResponseWriter,request *http.Request){
		err := handler(writer,request)
		if err != nil{
			log.Printf("Error handling request: %s",err.Error())
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFilelist))  //HandleFilelist 函数中包括了业务的逻辑，然后将所有异常抛出即可

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
