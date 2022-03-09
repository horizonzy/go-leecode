package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var slow = head
	var fast = head

	for fast != nil {
		slow = slow.Next

		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
