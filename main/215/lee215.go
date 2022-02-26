package main

import "fmt"

func main() {
	x := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	largest := findKthLargest(x, 3)
	fmt.Println(largest)
}

func findKthLargest(nums []int, k int) int {
	buildMaxHeap(nums)
	heapSize := len(nums) - 1
	for i := 0; i < k-1; i++ {
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		heapSize--
		heapify(nums, 0, heapSize)
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
	l := 2*i + 1
	r := 2*i + 2

	largest := i

	if l < size && nums[l] > nums[largest] {
		largest = l
	}
	if r < size && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, largest, size)
	}

}
