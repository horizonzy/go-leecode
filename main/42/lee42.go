package main

func main() {

}

func trap(height []int) int {

	//当前下标的雨水 = min(左边最大值, 右边最大值) - 当前下标的水位
	//通过动态规划计算每个下标对应的左边最大值和右边最大值
	//leftMax[i] = max(leftMax[i-1], nums[i])
	//rightMax[i] = max(rightMax[i+1], nums[i])

	len := len(height)

	var leftMax = make([]int, len)
	leftMax[0] = height[0]

	var rightMax = make([]int, len)
	rightMax[len-1] = height[len-1]

	for i := 1; i < len; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := len - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var result int

	for i := range height {
		result += min(leftMax[i], rightMax[i]) - height[i]
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
