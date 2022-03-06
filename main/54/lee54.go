package main

import "fmt"

func main() {
	var input = make([][]int, 2)
	input[0] = []int{2}
	input[1] = []int{3}
	//input[2] = []int{7, 8, 9}
	output := spiralOrder(input)
	for i := range output {
		fmt.Println(output[i])
	}
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	result = append(result, matrix[0][0])

	totalCount := len(matrix) * len(matrix[0])

	left := 0
	right := len(matrix[0]) - 1
	up := 1
	down := len(matrix) - 1
	i := 0
	j := 0

	direction := 0

	for totalCount != len(result) {
		switch direction {
		case 0:
			//右
			if j == right {
				direction = 1
				right--
			} else {
				j++
				result = append(result, matrix[i][j])
			}

		case 1:
			//下
			if i == down {
				direction = 2
				down--
			} else {
				i++
				result = append(result, matrix[i][j])
			}
		case 2:
			//左
			if j == left {
				direction = 3
				left++
			} else {
				j--
				result = append(result, matrix[i][j])
			}
		case 3:
			if i == up {
				direction = 0
				up++
			} else {
				//上
				i--
				result = append(result, matrix[i][j])
			}
		}
	}
	return result
}
