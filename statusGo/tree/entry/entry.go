package main

import (
	"fmt"
	"statusGo/tree"
)

//type myTreeNode struct {
//	node *tree.Node
//}

// Embedding 内嵌的方式扩展
type myTreeNode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Node.Print()
}

func main() {
	//var root tree.Node
	//	//root = tree.Node{Value: 3}
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Print("In-order traversal:")
	root.Traverse()
	fmt.Println()
	fmt.Print("My own post-order traversal:")
	root.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node){
		//node.Print()
		nodeCount ++
	})
	fmt.Println("Node count:",nodeCount)
}

//包
//如何扩充系统类型或者别人的类型
//1.定义别名，例子就是 queue
//2.使用组合，例子就是 postOrder() 使用了tree.Node 进行了扩充
//3.重载Embedding，重载可以直接将已有的方法直接使用，简化调用，并且可以继续加上自己的功能
