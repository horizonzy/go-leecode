package main

import "fmt"

func main() {
	x := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	largest := findKthLargest(x, 3)
	fmt.Println(largest)
}

func findKthLargest(nums []int, k int) int {
	quicksort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quicksort(nums []int, l, r int) {
	if l > r {
		return
	}

	base := nums[l]
	i := l
	j := r

	for i < j {
		for i < j {
			if nums[j] <= base {
				nums[i] = nums[j]
				i++
				break
			} else {
				j--
			}
		}

		for i < j {
			if nums[i] >= base {
				nums[j] = nums[i]
				j--
				break
			} else {
				i++
			}
		}
	}
	nums[i] = base
	quicksort(nums, l, i-1)
	quicksort(nums, i+1, r)
}
