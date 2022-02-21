package main

import "fmt"

func main() {
	x := "abcbc"
	size := lengthOfLongestSubstring(x)
	fmt.Println(size)
}

func lengthOfLongestSubstring(s string) int {
	var max int
	var indexMap = make(map[uint8]int)
	var l int
	for i := range s {
		sign := s[i]
		v, ok := indexMap[sign]
		if ok {
			len := len(indexMap)
			if len > max {
				max = len
			}
			for j := l; j < v; j++ {
				delete(indexMap, s[j])
			}
			indexMap[sign] = i
			l = v + 1
		} else {
			indexMap[sign] = i
		}
	}
	len := len(indexMap)
	if len > max {
		max = len
	}
	return max
}
