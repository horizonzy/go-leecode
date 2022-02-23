package main

import (
	"fmt"
	"math"
)

func main() {
	//grid = [[1, 3, 1], [1, 5, 1], [4, 2, 1]]

	var grid = make([][]int, 2)
	grid[0] = []int{1, 2, 3}
	grid[1] = []int{4, 5, 6}
	//grid[2] = []int{4, 2, 1}

	i := minPathSum(grid)
	fmt.Println(i)

}

func minPathSum(grid [][]int) int {
	//dp[i][j] = min(dp[i-1][j],min[i][j-1]) + grid[i][j]
	m := len(grid)
	n := len(grid[0])

	var dp = make([][]int, m)
	dp[0] = make([]int, n)
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
