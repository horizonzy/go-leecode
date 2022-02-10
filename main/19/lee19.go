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
	d := &ListNode{Val: 4}
	e := &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	result := removeNthFromEnd(a, 2)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	slow := head
	var preSlow *ListNode
	var index int
	for head != nil {
		index++
		head = head.Next
		if index > n {
			preSlow = slow
			slow = slow.Next
		}
	}
	if preSlow == nil {
		first = first.Next
		return first
	}
	preSlow.Next = slow.Next
	return first
}
