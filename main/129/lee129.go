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
	a.Left = b
	a.Right = c
	rs := sumNumbers(a)
	fmt.Println(rs)
}

var numbers []int

func sumNumbers(root *TreeNode) int {
	numbers = []int{}
	backtrack(root.Val, root)
	var res int

	for i := range numbers {
		res += numbers[i]
	}
	return res
}

func backtrack(target int, node *TreeNode) {
	if node.Left == nil && node.Right == nil {
		numbers = append(numbers, target)
	}
	if node.Left != nil {
		backtrack(target*10+node.Left.Val, node.Left)
	}
	if node.Right != nil {
		backtrack(target*10+node.Right.Val, node.Right)
	}

}
