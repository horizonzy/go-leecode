package main

import "fmt"

func main() {
	//[] 2，10，7
	var input = []int{0}
	res := lengthOfLIS(input)
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {
	var result = 1
	//dp[i] = max(dp[j] + 1)
	var dp = make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		var max int
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > max {
					max = dp[j] + 1
				}
			}
		}
		if max > dp[i] {
			dp[i] = max
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
