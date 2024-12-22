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

func isSameTree(p *TreeNode, q *TreeNode) bool {

	l := true
	r := true

	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	l = isSameTree(p.Left, q.Left)
	r = isSameTree(p.Right, q.Right)

	return l && r
}

func main() {

	n4 := NewTreeNode(9, nil, nil)
	n3 := NewTreeNode(7, nil, nil)
	n2 := NewTreeNode(15, nil, nil)
	n1 := NewTreeNode(20, n2, n3)
	root1 := NewTreeNode(3, n4, n1)

	n42 := NewTreeNode(9, nil, nil)
	n32 := NewTreeNode(7, nil, nil)
	n22 := NewTreeNode(15, nil, nil)
	n12 := NewTreeNode(21, n22, n32)
	root2 := NewTreeNode(3, n42, n12)

	fmt.Println(isSameTree(root1, root2))
}
