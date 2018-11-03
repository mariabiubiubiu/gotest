package main

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (t *TreeNode) Print() {
	fmt.Println(t.Value, "***")
}

func (t *TreeNode) Traverse() {
	if t != nil {
		t.Left.Traverse()
		t.Print()
		t.Right.Traverse()
	}
}

func main() {
	var root TreeNode
	root.Value = 1
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{Value: 4}
	root.Traverse()
}
