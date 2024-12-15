package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	value int
	next  *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		value: val,
		next:  nil,
	}
}

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]*ListNode)

	current := head
	for current != nil {

		if _, e := visited[current]; e {
			return true
		} else {
			visited[current] = current
		}

		current = current.next
	}

	return false
}

func main() {

	head := NewListNode(3)
	v1 := NewListNode(2)
	v2 := NewListNode(0)
	v3 := NewListNode(-4)

	head.next = v1
	v1.next = v2
	v2.next = v3
	v3.next = v1

	fmt.Println(hasCycle(head))
}
