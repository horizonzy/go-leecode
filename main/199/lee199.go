package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := &TreeNode{Val: 1}
	b := &TreeNode{Val: 2}
	c := &TreeNode{Val: 3}
	d := &TreeNode{Val: 4}
	e := &TreeNode{Val: 5}

	a.Left = b
	a.Right = c
	b.Right = e
	c.Right = d

	result := rightSideView(a)
	fmt.Printf("%v", result)

}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	var result []int
	queue = append(queue, root)
	var currentTmp []*TreeNode
	for {
		node := queue[0]
		queue = queue[1:]
		if len(queue) == 0 {
			result = append(result, node.Val)
		}
		if node.Left != nil {
			currentTmp = append(currentTmp, node.Left)
		}
		if node.Right != nil {
			currentTmp = append(currentTmp, node.Right)
		}
		if len(queue) == 0 {
			if len(currentTmp) == 0 {
				break
			} else {
				queue = append(queue, currentTmp...)
				currentTmp = []*TreeNode{}
			}
		}
	}
	return result
}
