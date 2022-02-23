package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	restoreIpAddresses("1012322102")
}

func restoreIpAddresses(s string) []string {
	len := len(s)
	var result []string
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				if i+j+k >= len {
					continue
				}
				s1 := s[:i]
				s2 := s[i : i+j]
				s3 := s[i+j : i+j+k]
				s4 := s[i+j+k:]
				if checkStrValid(s1) && checkStrValid(s2) && checkStrValid(s3) && checkStrValid(s4) {
					result = append(result, fmt.Sprintf("%s.%s.%s.%s", s1, s2, s3, s4))
				}
			}
		}
	}
	return result
}

func checkStrValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) > 1 && strings.HasPrefix(s, "0") {
		return false
	}
	if v, _ := strconv.Atoi(s); v > 255 {
		return false
	}
	return true
}
