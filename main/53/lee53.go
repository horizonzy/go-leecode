package main

func main() {
	x := []int{1}
	resutl := maxSubArray(x)
	println(resutl)
}

func maxSubArray(nums []int) int {
	//dp[i] = max(dp[i-1]+nums[i], nums[i])
	var result int
	len := len(nums)

	var dp = make([]int, len)
	dp[0] = nums[0]
	result = nums[0]

	for i := 1; i < len; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
