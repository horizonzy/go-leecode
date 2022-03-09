package main

import "fmt"

func main() {
	res := firstMissingPositive([]int{3, 4, -1, 1})
	fmt.Println(res)
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] > n || nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := range nums {
		v := abs(nums[i])
		if v > 0 && v <= n {
			nums[v-1] = -abs(nums[v-1])
		}
	}
	for i := range nums {
		v := nums[i]
		if v > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
