package main

import "fmt"

func main() {
	input := []int{2, 7, 11, 15}
	sum := twoSum(input, 18)
	for i := range sum {
		fmt.Printf("%d ", sum[i])
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		expect := target - nums[i]
		if v, ok := m[expect]; ok {
			return []int{v, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
