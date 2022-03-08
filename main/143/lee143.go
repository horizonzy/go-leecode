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
	//e := &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	//d.Next = e
	reorderList(a)
	fmt.Println("aaa")
}

func reorderList(head *ListNode) {
	first := head

	var slowNode = head
	for head != nil && head.Next != nil {
		slowNode = slowNode.Next
		head = head.Next.Next
	}

	var middle = slowNode
	var pre *ListNode

	for slowNode != nil {
		tmp := slowNode.Next
		slowNode.Next = pre
		pre = slowNode
		slowNode = tmp
	}

	for first != middle {
		tmp := first.Next
		tmp1 := pre.Next

		first.Next = pre
		pre.Next = tmp

		pre = tmp1
		first = tmp
	}
	first.Next = nil
}
