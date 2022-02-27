package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{0, 0, 0, 0}

	sum := threeSum(input)

	for _, ele := range sum {
		for _, eleb := range ele {
			fmt.Printf("%d ", eleb)
		}
		fmt.Println("")
	}
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	len := len(nums)

	for i := 0; i < len-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		r := len - 1
		for j := i + 1; j < len-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1] > 0 {
				break
			}
			c := r
			for k := c; k > j; k-- {
				if k < c && nums[k] == nums[k+1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] < 0 {
					break
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					r = k - 1
				}
			}
		}
	}
	return result
}
