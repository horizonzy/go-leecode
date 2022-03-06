package main

func main() {
	var input = []int{1, 3}
	search(input, 3)
}

func search(nums []int, target int) int {
	return midSearch(0, len(nums)-1, target, nums)
}

func midSearch(l, r, target int, nums []int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[l] <= nums[mid] {
		if nums[l] <= target && target <= nums[mid] {
			return midSearch(l, mid, target, nums)
		} else {
			return midSearch(mid+1, r, target, nums)
		}
	} else {
		if nums[mid] <= target && target <= nums[r] {
			return midSearch(mid, r, target, nums)
		} else {
			return midSearch(l, mid-1, target, nums)
		}
	}
}
