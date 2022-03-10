package main

import "fmt"

func main() {
	s := climbStairs(5)
	fmt.Println(s)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	//dp[i] = dp[i-1] + dp[i-2]

	var dp = make([]int, n+1)

	dp[2] = 2
	dp[1] = 1
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
