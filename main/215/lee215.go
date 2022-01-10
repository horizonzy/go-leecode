package main

import "fmt"

func main() {
	x := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	largest := findKthLargest(x, 3)
	fmt.Println(largest)
}

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := 0; i < k-1; i++ {
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		heapSize--
		heapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, size int) {
	for i := size / 2; i >= 0; i-- {
		heapify(a, i, size)
	}
}

func heapify(a []int, i, size int) {
	l := i*2 + 1
	r := i*2 + 2
	largest := i

	if l < size && a[l] > a[largest] {
		largest = l
	}
	if r < size && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		heapify(a, largest, size)
	}
}
