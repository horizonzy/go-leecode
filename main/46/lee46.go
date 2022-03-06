package main

import "fmt"

func main() {
	var input = []int{1, 2, 3}
	output := permute(input)
	fmt.Printf("%+v", output)
}

var total [][]int

func permute(nums []int) [][]int {
	total = [][]int{}
	var result []int
	backtrack(result, nums)
	return total
}

func backtrack(result, nums []int) {
	if nil == nums {
		total = append(total, result)
		return
	}
	if len(nums) == 1 {
		tmp1 := append(result, nums[0])
		backtrack(tmp1, nil)
	} else {
		for i := range nums {
			tmp1 := append(result, nums[i])

			n1 := nums[0:i]
			n2 := nums[i+1:]
			var tmp2 []int
			tmp2 = append(tmp2, n1...)
			tmp2 = append(tmp2, n2...)
			backtrack(tmp1, tmp2)
		}
	}
}
