package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	//c := &ListNode{Val: 3}
	//d := &ListNode{Val: 4}
	//e := &ListNode{Val: 5}
	a.Next = b
	//b.Next = c
	//c.Next = d
	//d.Next = e
	result := removeNthFromEnd(a, 1)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := head
	var pre *ListNode
	var slow = head
	var index int
	for head != nil {
		if index > n-1 {
			pre = slow
			slow = slow.Next
		}
		index++
		head = head.Next
	}
	if pre != nil {
		pre.Next = slow.Next
	} else {
		result = result.Next
	}
	return result
}
