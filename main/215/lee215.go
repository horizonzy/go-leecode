package main

import "fmt"

func main() {
	x := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	largest := findKthLargest(x, 3)
	fmt.Println(largest)
}

func findKthLargest(nums []int, k int) int {
	buildMaxHeap(nums)
	size := len(nums) - 1
	for i := 0; i < k-1; i++ {
		nums[0], nums[size] = nums[size], nums[0]
		size--
		heapify(nums, 0, size)
	}
	return nums[0]
}

func buildMaxHeap(nums []int) {
	len := len(nums)

	for i := len / 2; i >= 0; i-- {
		heapify(nums, i, len)
	}
}

func heapify(nums []int, i, size int) {
	l := i*2 + 1
	r := l + 1
	var max = i

	if l < size && nums[l] > nums[max] {
		max = l
	}

	if r < size && nums[r] > nums[max] {
		max = r
	}

	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heapify(nums, max, size)
	}
}
