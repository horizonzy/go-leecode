package main

import "fmt"

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func main() {
	m := letterCombinations("23")
	for i := range m {
		fmt.Println(m[i])
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	combinations = []string{}
	backtrack(digits, "", 0)
	return combinations
}

func backtrack(digits, result string, level int) {
	if len(digits) == level {
		combinations = append(combinations, result)
		return
	}
	chs := phoneMap[string(digits[level])]
	for i := range chs {
		backtrack(digits, result+string(chs[i]), level+1)
	}
}
