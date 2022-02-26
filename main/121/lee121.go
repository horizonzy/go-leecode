package main

func main() {

	var input = []int{5, 4, 3, 2, 1}
	result := maxProfit(input)
	println(result)

}

func maxProfit(prices []int) int {
	//dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	//dp[i][1] = max(-prices[i], dp[i-1][1])

	len := len(prices)

	var dp = make([][]int, len)

	for i := 0; i < len; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(-prices[i], dp[i-1][1])
	}
	return dp[len-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
