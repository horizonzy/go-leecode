package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for slow != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
	}
	return false
}
