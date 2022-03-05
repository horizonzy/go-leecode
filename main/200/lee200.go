package main

import "fmt"

func main() {
	m := make([][]byte, 3)
	m[0] = []byte{'1', '1', '1'}
	m[1] = []byte{'0', '1', '0'}
	m[2] = []byte{'1', '1', '1'}
	result := numIslands(m)
	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	var m = make(map[string]struct{})
	len1 := len(grid)
	len2 := len(grid[0])

	var count int
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if dsf(grid, i, j, m) {
				count++
			}
		}
	}
	return count
}

func dsf(grid [][]byte, i, j int, m map[string]struct{}) bool {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 {
		return false
	}
	if grid[i][j] == '0' {
		return false
	}
	key := buildKey(i, j)
	if _, ok := m[key]; !ok {
		m[key] = struct{}{}
		//上
		dsf(grid, i-1, j, m)
		//左
		dsf(grid, i, j-1, m)
		// 右
		dsf(grid, i, j+1, m)
		// 下
		dsf(grid, i+1, j, m)

		return true
	} else {
		return false
	}
}

func buildKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}
