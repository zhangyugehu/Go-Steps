package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}
// 为结构定义方法
func (node Node) Print(){
	fmt.Println(node.Value)
}
func (node *Node) SetValue(value int){
	if node == nil {
		return
	}
	node.Value = value
}

func CreateTreeNode(value int) *Node {
	return &Node{Value:value}
}

func main() {
	var root Node

	root = Node{Value:3}
	root.Left =&Node{}
	root.Right =&Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = CreateTreeNode(4)

	root.Left.SetValue(6)
	root.Left.Print()

	root.Print()

	// nil也可以调用方法
	root.Right.Left.Left.Print()


	nodes:=[]Node{
		{Value:3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
}