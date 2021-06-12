package main

import "fmt"

func main() {
	c := calcu("pwwkew")
	fmt.Println(c)
}

func calcu(s string) int {
	if len(s) == 0 {
		return 00
	}

	m := make(map[int32]int)

	maxLength := -1
	currentLeng := -1

	for index, value := range s {

		if maxLength == -1 {
			maxLength = 1
			currentLeng = 1
			m[value] = index
			continue
		}

		if beforeIndex, ok := m[value]; !ok {
			if maxLength == currentLeng {
				maxLength++
				currentLeng++
			} else {
				currentLeng++
			}
			if currentLeng > maxLength {
				maxLength = currentLeng
			}
			m[value] = index
		} else {
			c := s[beforeIndex+1 : index+1]
			m = map[int32]int{}
			currentLeng = index - beforeIndex
			for i, v := range c {
				m[v] = i + beforeIndex + 1
			}
		}
	}
	return maxLength
}
