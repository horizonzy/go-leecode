package main

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

	v := zigzagLevelOrder(a)
	println(v)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int

	var data []*TreeNode
	data = append(data, root)
	level := make(map[*TreeNode]int)
	i := 0
	level[root] = 0
	var currentLevelResult []int
	var reverse bool
	for len(data) != 0 {
		current := data[0]
		data = data[1:]
		currentLevel := level[current]
		if current.Left != nil {
			level[current.Left] = currentLevel + 1
			data = append(data, current.Left)
		}
		if current.Right != nil {
			level[current.Right] = currentLevel + 1
			data = append(data, current.Right)
		}

		if currentLevel == i {
			currentLevelResult = append(currentLevelResult, current.Val)
		}
		if currentLevel > i {
			i = currentLevel
			if reverse {
				for i, n := 0, len(currentLevelResult); i < n/2; i++ {
					currentLevelResult[i], currentLevelResult[n-i-1] = currentLevelResult[n-i-1], currentLevelResult[i]
				}
				result = append(result, currentLevelResult)
				reverse = !reverse
			} else {
				result = append(result, currentLevelResult)
				reverse = !reverse
			}
			currentLevelResult = []int{}
			currentLevelResult = append(currentLevelResult, current.Val)
		}
	}
	if reverse {
		for i, n := 0, len(currentLevelResult); i < n/2; i++ {
			currentLevelResult[i], currentLevelResult[n-i-1] = currentLevelResult[n-i-1], currentLevelResult[i]
		}
		result = append(result, currentLevelResult)
		reverse = !reverse
	} else {
		result = append(result, currentLevelResult)
		reverse = !reverse
	}
	return result
}
