package main

import "fmt"

func main() {
	c := []int{5, 1, 1}
	nextPermutation(c)
	fmt.Printf("%v", c)
}

func nextPermutation(nums []int) {
	len := len(nums)
	if len == 0 || len == 1 {
		return
	}
	current := len - 2
	var hasChange bool
	for current >= 0 {
		if nums[current] >= nums[current+1] {
			current--
		} else {
			hasChange = true
			var min = current + 1
			for i := min + 1; i < len; i++ {
				if nums[i] < nums[min] && nums[i] > nums[current] {
					min = i
				}
			}
			nums[current], nums[min] = nums[min], nums[current]

			tmp := current + 1
			for tmp <= len-2 {
				for i := tmp + 1; i < len; i++ {
					if nums[tmp] > nums[i] {
						nums[tmp], nums[i] = nums[i], nums[tmp]
					}
				}
				tmp++
			}
			break
		}
	}
	if !hasChange {
		for i := 0; i < len/2; i++ {
			nums[i], nums[len-1-i] = nums[len-1-i], nums[i]
		}
	}
}
