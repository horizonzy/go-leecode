package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := restoreIpAddresses("0000")
	for i := range result {
		fmt.Println(result[i])
	}
}

var total []string

func restoreIpAddresses(s string) []string {
	if s == "" {
		return nil
	}
	total = []string{}
	backtrack(0, 0, "", s)
	return total
}

func backtrack(start, level int, current string, source string) {
	if level == 4 {
		if start == len(source) {
			total = append(total, current)
		}
		return
	}

	for i := 1; i <= 3; i++ {
		if start+i > len(source) {
			break
		}

		child := source[start : start+i]
		if len(child) > 1 && child[0] == '0' {
			break
		}
		if v, _ := strconv.Atoi(child); v > 255 {
			continue
		}
		if len(current) == 0 {
			backtrack(start+i, level+1, child, source)
		} else {
			backtrack(start+i, level+1, current+"."+child, source)
		}
	}
}
