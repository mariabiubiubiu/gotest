package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}

func (node *Node)GetValue()  {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int)error {
	if node == nil{
		return errors.New("node is null")
	}
	node.Value = value
	return nil
}

func (node *Node) PreOrder(){
	if node == nil{
		return
	}
	node.GetValue()
	node.Left.PreOrder()
	node.Right.PreOrder()
}

//中序遍历
func (node *Node) middleOrder() {
	if node == nil {
		return
	}
	node.Left.middleOrder()
	node.GetValue()
	node.Right.middleOrder()
}

//后序遍历
func (node *Node) postOrder() {
	if node == nil {
		return
	}
	node.Left.postOrder()
	node.Right.postOrder()
	node.GetValue()
}

//创建结点
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func main() {
	root := Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Right.Left.SetValue(4)
	root.Left.Right = CreateNode(2)

	fmt.Print("前序遍历: ")
	root.PreOrder()
	fmt.Println()
	fmt.Print("中序遍历: ")
	root.middleOrder()
	fmt.Println()
	fmt.Print("后序遍历: ")
	root.postOrder()
}
