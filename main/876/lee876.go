package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	a.Next = b
	b.Next = c
	c.Next = d

	middleNode(a)

}

func middleNode(head *ListNode) *ListNode {
	var slowNode = head
	for head != nil && head.Next != nil {
		slowNode = slowNode.Next
		head = head.Next.Next
	}
	return slowNode
}
