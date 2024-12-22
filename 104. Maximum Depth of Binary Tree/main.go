package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int, l *TreeNode, r *TreeNode) *TreeNode {

	return &TreeNode{
		Val:   v,
		Left:  l,
		Right: r,
	}

}

func deep(node *TreeNode, n int) int {

	if node == nil {
		return n
	}

	return max(deep(node.Left, n+1), deep(node.Right, n+1))
}

func maxDepth(root *TreeNode) int {
	return deep(root, 0)
}

func main() {

	n4 := NewTreeNode(9, nil, nil)
	n3 := NewTreeNode(7, nil, nil)
	n2 := NewTreeNode(15, nil, nil)
	n1 := NewTreeNode(20, n2, n3)
	root := NewTreeNode(3, n4, n1)

	fmt.Println(maxDepth(root))
}
