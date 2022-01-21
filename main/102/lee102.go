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
	var level int
	var queue []*TreeNode
	levelSize := make(map[int]int)
	levelSize[0] = 1
	queue = append(queue, root)
	for len(queue) > 0 {
		size := levelSize[level]
		var currentSize int
		var childResult []int
		for i := 0; i < size; i++ {
			tmp := queue[0]
			childResult = append(childResult, tmp.Val)
			queue = queue[1:]
			if tmp.Left != nil {
				currentSize++
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				currentSize++
				queue = append(queue, tmp.Right)
			}
		}
		result = append(result, childResult)
		level++
		levelSize[level] = currentSize
	}
	return result
}

//func levelOrder(root *TreeNode) [][]int {
//	var result []*[]int
//
//	if root == nil {
//		return [][]int{}
//	}
//	base := &[]int{root.Val}
//	result = append(result, base)
//
//	transfer(root, 1, &result)
//
//	var x [][]int
//	for i := range result {
//		if len(*result[i]) != 0 {
//			x = append(x, *result[i])
//		}
//	}
//	return x
//}
//
//func transfer(root *TreeNode, level int, array *[]*[]int) {
//	var current *[]int
//
//	if len(*array) > level {
//		current = (*array)[level]
//	} else {
//		current = &[]int{}
//		*array = append(*array, current)
//	}
//
//	if root.Left != nil {
//		*current = append(*current, root.Left.Val)
//		transfer(root.Left, level+1, array)
//	}
//	if root.Right != nil {
//		*current = append(*current, root.Right.Val)
//		transfer(root.Right, level+1, array)
//	}
//}
