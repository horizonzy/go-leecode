package main

import "fmt"

func main() {
	x := "abcbc"
	size := lengthOfLongestSubstring(x)
	fmt.Println(size)
}

func lengthOfLongestSubstring(s string) int {

	size := len(s)

	var result int
	m := map[uint8]int{}
	var l int

	for i := 0; i < size; i++ {
		if v, ok := m[s[i]]; !ok {
			m[s[i]] = i
		} else {
			if i-l > result {
				result = i - l
			}
			for j := l; j < v; j++ {
				delete(m, s[j])
			}
			m[s[i]] = i
			l = v + 1
		}
	}
	if size-l > result {
		result = size - l
	}

	return result
}
