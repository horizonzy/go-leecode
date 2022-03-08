package main

func main() {
	a := &TreeNode{Val: 3}
	b := &TreeNode{Val: 9}
	c := &TreeNode{Val: 20}
	d := &TreeNode{Val: 15}
	e := &TreeNode{Val: 7}

	a.Left = b
	a.Right = c
	c.Left = d
	c.Right = e

	order := levelOrder(a)

	println(order)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int

	result = append(result, []int{root.Val})

	var queue = []*TreeNode{root}

	var tmp []*TreeNode
	var tmpResult []int

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left != nil {
			tmp = append(tmp, current.Left)
			tmpResult = append(tmpResult, current.Left.Val)
		}

		if current.Right != nil {
			tmp = append(tmp, current.Right)
			tmpResult = append(tmpResult, current.Right.Val)
		}

		if len(queue) == 0 && len(tmp) > 0 {
			queue = append(queue, tmp...)
			result = append(result, tmpResult)
			tmp = []*TreeNode{}
			tmpResult = []int{}
		}
	}
	return result
}
