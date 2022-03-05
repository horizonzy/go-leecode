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
	c.Left = d
	d.Right = e

	result := lowestCommonAncestor(a, c, e)
	fmt.Println(result.Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := make(map[*TreeNode]*TreeNode)
	traverse(root, parent)
	var parent1 []*TreeNode
	for tmp := p; tmp != nil; {
		parent1 = append(parent1, tmp)
		tmp = parent[tmp]
	}
	var parent2 []*TreeNode
	for tmp := q; tmp != nil; {
		parent2 = append(parent2, tmp)
		tmp = parent[tmp]
	}

	for i := range parent1 {
		for i2 := range parent2 {
			if parent2[i2] == parent1[i] {
				return parent2[i2]
			}
		}
	}
	return nil
}

func traverse(node *TreeNode, parent map[*TreeNode]*TreeNode) {
	if node.Left != nil {
		parent[node.Left] = node
		traverse(node.Left, parent)
	}
	if node.Right != nil {
		parent[node.Right] = node
		traverse(node.Right, parent)
	}
}
