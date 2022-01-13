package main

import "fmt"

func main() {
	x := []int{1}
	array := maxSubArray(x)
	fmt.Println(array)

}

//动态规划
//func maxSubArray(nums []int) int {
//	var max = nums[0]
//	size := len(nums)
//	for i := 1; i < size; i++ {
//		if nums[i-1] > 0 {
//			nums[i] += nums[i-1]
//		}
//		if nums[i] > max {
//			max = nums[i]
//		}
//	}
//	return max
//}

//贪心
func maxSubArray(nums []int) int {
	var max = -10001
	var sum int
	for i := range nums {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
