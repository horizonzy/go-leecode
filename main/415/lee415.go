package main

import (
	"fmt"
	"strconv"
)

func main() {

	res := addStrings("123456789", "987654321")
	fmt.Println(res)
}

func addStrings(num1 string, num2 string) string {

	var result string

	var add int
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x, _ = strconv.Atoi(string(num1[i]))
		}
		if j >= 0 {
			y, _ = strconv.Atoi(string(num2[j]))
		}
		sum := x + y + add
		result = strconv.Itoa(sum%10) + result
		add = sum / 10
	}
	return result
}
