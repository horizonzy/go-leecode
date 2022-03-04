package main

import "fmt"

func main() {
	result := generateParenthesis(3)
	for i := range result {
		fmt.Println(result[i])
	}
}

var total []string

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	total = []string{}
	backtrack(0, 0, n, "")
	return total
}

func backtrack(left, right, max int, target string) {
	if len(target) == max*2 {
		total = append(total, target)
		return
	}
	if left < max {
		backtrack(left+1, right, max, target+"(")
	}
	if right < left {
		backtrack(left, right+1, max, target+")")
	}
}
