package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}

	a.Next = b
	b.Next = c
	result := reverseList(a)

	for ; result.Next != nil; result = result.Next {
		fmt.Println(result.Val)
	}
	fmt.Println(result.Val)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}
	return prev
}
