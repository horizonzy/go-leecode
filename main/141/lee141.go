package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var slow = head
	var fast = head

	for slow != nil {
		slow = slow.Next

		if fast == nil || fast.Next == nil {
			return false
		} else {
			fast = fast.Next.Next
		}

		if slow == fast {
			return true
		}
	}
	return false
}
