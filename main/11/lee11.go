package main

func main() {

}

func maxArea(height []int) int {
	var max int
	l := 0
	r := len(height) - 1
	for l < r {
		lv := height[l]
		rv := height[r]
		min := min(lv, rv)
		current := (r - l) * min
		if current > max {
			max = current
		}

		if min == lv {
			l++
		} else {
			r--
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
