package main

import "fmt"

func main() {
	result := isValid("(){}}{")
	fmt.Println(result)
}

func isValid(s string) bool {
	var before []byte
	for i := range s {
		if s[i] == ']' {
			if before == nil || len(before) == 0 {
				return false
			}
			c := before[len(before)-1]
			before = before[:len(before)-1]
			if c != '[' {
				return false
			}
		} else if s[i] == '}' {
			if before == nil || len(before) == 0 {
				return false
			}
			c := before[len(before)-1]
			before = before[:len(before)-1]
			if c != '{' {
				return false
			}
		} else if s[i] == ')' {
			if before == nil || len(before) == 0 {
				return false
			}
			c := before[len(before)-1]
			before = before[:len(before)-1]
			if c != '(' {
				return false
			}
		} else {
			before = append(before, s[i])
		}
	}
	return len(before) == 0
}
