package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	resutl := reverseKGroup(node1, 3)
	for resutl != nil {
		fmt.Printf("%d ", resutl.Val)
		resutl = resutl.Next
	}

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	result := head
	var first = true

	var last *ListNode
	var prelast *ListNode

	var prev *ListNode
	var i int
	for head != nil {
		if i == 0 {
			last = head
		}
		i++
		tmp := head.Next
		head.Next = prev
		if i == k {
			if first {
				result = head
				first = false
			}
			if prelast != nil {
				prelast.Next = head
			}
			prelast = last
			prev = nil
			i = 0
		} else {
			if tmp == nil {
				var pre *ListNode
				for head != nil {
					tmp1 := head.Next
					head.Next = pre
					pre = head
					head = tmp1
				}
				prelast.Next = pre
			}
			prev = head
		}
		head = tmp
	}
	return result
}
