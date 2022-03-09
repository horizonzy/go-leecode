package main

import "fmt"

func main() {
	c := mySqrt(8)
	fmt.Println(c)
}

func mySqrt(x int) int {
	l, r := 0, x
	var rs int
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			rs = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return rs
}
