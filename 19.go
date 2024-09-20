package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Val: 0, Next: head}
	dummy := res

	for i := 0; i < n; i++ {
		head = head.Next
	}

	for head != nil {
		head = head.Next
		dummy = dummy.Next
	}

	dummy.Next = dummy.Next.Next
	return res.Next
}

func walkAndPrint(head *ListNode) {
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	removeNthFromEnd(node1, 1)
	walkAndPrint(node1)
}
