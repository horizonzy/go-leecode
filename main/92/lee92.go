package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//a := &ListNode{Val: 1}
	//b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	//d := &ListNode{Val: 4}
	e := &ListNode{Val: 5}
	//a.Next = b
	//b.Next = c
	//c.Next = d
	//d.Next = e
	c.Next = e
	reverseBetween(c, 1, 2)

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var first = head

	var leftPreNode *ListNode

	var leftNode *ListNode

	var index int

	var pre *ListNode

	for head != nil {
		index++
		if index == left-1 {
			leftPreNode = head
		}

		if index >= left && index <= right {
			if index == left {
				leftNode = head
			}
			tmp := head.Next
			head.Next = pre
			pre = head
			head = tmp
			if index == right {
				if leftPreNode == nil {
					first = pre
					leftNode.Next = head
				} else {
					leftNode.Next = head
					leftPreNode.Next = pre
				}
			}
			continue
		}
		head = head.Next
	}
	return first
}
