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
	var leveMap = make(map[*TreeNode]int)
	var queue []*TreeNode
	queue = append(queue, root)
	leveMap[root] = 0
	i := 0
	var tmp []int
	for len(queue) != 0 {
		current := queue[0]
		currentLevel := leveMap[current]
		queue = queue[1:]
		if current.Left != nil {
			leveMap[current.Left] = currentLevel + 1
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			leveMap[current.Right] = currentLevel + 1
			queue = append(queue, current.Right)
		}
		if currentLevel == i {
			tmp = append(tmp, current.Val)
		}
		if currentLevel > i {
			i = currentLevel
			result = append(result, tmp)
			tmp = []int{}
			tmp = append(tmp, current.Val)
		}
	}
	result = append(result, tmp)
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
