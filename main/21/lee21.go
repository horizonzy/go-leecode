package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 4}
	a.Next = b
	b.Next = c

	d := &ListNode{Val: 1}
	e := &ListNode{Val: 3}
	f := &ListNode{Val: 4}
	d.Next = e
	e.Next = f

	new := mergeTwoLists(a, d)

	for new != nil {
		fmt.Printf("%d ", new.Val)
		new = new.Next
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	dummy := result

	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 == nil {
			dummy.Next = list2
			dummy = list2

			list2 = list2.Next
			continue

		}
		if list2 == nil {
			dummy.Next = list1
			dummy = list1

			list1 = list1.Next
			continue
		}

		if list1.Val > list2.Val {
			dummy.Next = list2

			tmp2 := list2.Next

			dummy = list2

			list2 = tmp2
		} else {
			dummy.Next = list1

			tmp1 := list1.Next

			dummy = list1

			list1 = tmp1

		}
	}
	return result.Next
}
