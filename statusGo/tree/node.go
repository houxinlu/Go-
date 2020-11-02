package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//使用指针作为方法接受者
//只有使用指针才可以改变结构内容，指针实现引用传递，不然拷贝值传递了，本身值不会进行改变
//nil指针也可以调用方法
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

//使用工厂函数，避免使用new() 增加可扩展性
//注意返回了局部变量的地址
//堆？栈？不用关系，GO语言有垃圾回收器
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//值接受者vs指针接受者
//要改变内容必须使用指针接受者
//结构过大也考虑使用指针接收者
//一致性：如有指针接受者，最好都是指针接收者
//值接受者是go语言特有的
//值/指针接收者均可接受值/指针


//封装
//名字一般使用CamelCase
//首字母大写:public  公共
//首字母小写:private 私有
//
//包
//每个目录一个包
//main包 包含可执行入口
//为结构定义的方法必须放在同一个包内
//可以是不同文件


