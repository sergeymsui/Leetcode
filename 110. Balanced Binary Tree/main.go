package main

import (
	"fmt"
)

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

func deep(root *TreeNode, n int) (bool, int) {

	if root == nil {
		return true, n
	}

	l, deepLeft := deep(root.Left, n)
	r, deepRight := deep(root.Right, n)

	return (l && r && ((deepLeft-deepRight)*(deepLeft-deepRight)) <= 1), max(deepLeft, deepRight) + 1
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	ans, _ := deep(root, 0)

	return ans
}

func main() {

	n4 := NewTreeNode(7, nil, nil)
	n3 := NewTreeNode(15, nil, nil)

	n2 := NewTreeNode(20, n3, n4)
	n1 := NewTreeNode(9, nil, nil)

	root := NewTreeNode(3, n1, n2)

	fmt.Println(isBalanced(root))
}
