package main

import "fmt"

/**
 * Definition for singly-linked list.
*
*/
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

	resutl := reverseKGroup(node1, 2)
	println(resutl)

}

func print(head *ListNode) {
	currentNode := head
	for {
		if currentNode.Next == nil {
			fmt.Println(currentNode.Val)
			break
		}
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	tmp := []*ListNode{}

	var result *ListNode
	var beforeLast *ListNode
	currentNode := head
	i := 0

	for {
		if currentNode.Next == nil {
			i++
			tmp = append(tmp, currentNode)
			if i == k {
				start, last := reverse(tmp)
				if result == nil {
					result = start
				}
				if beforeLast != nil {
					beforeLast.Next = start
				}
				beforeLast = last

				tmp = []*ListNode{}
				i = 0
			} else {
				if beforeLast == nil {
					result = tmp[0]
				} else {
					beforeLast.Next = tmp[0]
				}
			}
			break
		}
		tmp = append(tmp, currentNode)
		currentNode = currentNode.Next
		i++
		if i == k {
			start, last := reverse(tmp)
			if result == nil {
				result = start
			}

			if beforeLast != nil {
				beforeLast.Next = start
			}
			beforeLast = last
			tmp = []*ListNode{}
			i = 0
		}
	}
	return result
}

func reverse(list []*ListNode) (*ListNode, *ListNode) {
	var result *ListNode
	var last *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		if i == 0 {
			list[i].Next = nil
			last = list[i]
		} else {
			if result == nil {
				result = list[i]
			}
			list[i].Next = list[i-1]
		}
	}
	return result, last
}
