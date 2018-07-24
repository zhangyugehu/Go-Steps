package main

import (
	"fmt"
	"study/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left =&tree.Node{}
	root.Right =&tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(4)

	root.Left.SetValue(6)
	root.Left.Print()

	root.Print()

	// nil也可以调用方法
	//root.Right.Left.Left.Print()


	nodes:=[]tree.Node{
		{Value:3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
}