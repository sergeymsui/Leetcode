package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) int {

	leftMax := root.value
	rightMax := root.value

	if root.left != nil {
		leftMax = Solution(root.left)
	}

	if root.right != nil {
		rightMax = Solution(root.right)
	}

	return max(root.value, leftMax, rightMax)
}

func main() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, &node3, nil}
	if s := Solution(&node4); s != 3 {
		panic("WA")
	} else {
		fmt.Println(s)
	}
}
