package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}

	sum := threeSum(input)

	for _, ele := range sum {
		for _, eleb := range ele {
			fmt.Printf("%d ", eleb)
		}
		fmt.Println("")
	}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	size := len(nums)

	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := size - 1
		target := -1 * nums[i]
		for j := i + 1; j < size-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for r > j && nums[r]+nums[j] > target {
				r--
			}
			if j == r {
				break
			}

			if nums[r]+nums[j] == target {
				result = append(result, []int{nums[i], nums[j], nums[r]})
			}
		}
	}
	return result
}
