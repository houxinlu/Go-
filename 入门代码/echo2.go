// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

//func main() {
//	fmt.Println("hello world")
//}

//func main() {
//	var s,sep string
//	for i:=1;i<len(os.Args);i++{
//		s += sep + os.Args[i]
//		sep = ""
//	}
//	fmt.Println(s)
//}

func main() {
//	s, sep := "", ""
	for temp,arg := range os.Args[1:]{
//		s += sep + arg
//		sep = ""
            fmt.Println(temp,arg)
	}
}
