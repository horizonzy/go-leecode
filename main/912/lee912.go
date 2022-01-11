package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 1}
	x := sortArray(nums)
	for _, ele := range x {
		fmt.Printf("%d ", ele)
	}
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l < r {
		i, j := l, r
		base := nums[l]

		for i < j {
			for ; i < j; j-- {
				if nums[j] <= base {
					nums[i] = nums[j]
					i++
					break
				}
			}

			for ; i < j; i++ {
				if nums[i] >= base {
					nums[j] = nums[i]
					j--
					break
				}
			}
		}
		nums[i] = base
		quickSort(nums, l, i-1)
		quickSort(nums, i+1, r)
	}
}
