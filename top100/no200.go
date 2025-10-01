package top100

import "fmt"

func numIslands(grid [][]byte) int {
	ret := 0
	d := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			for _, t := range d {
				dfs(grid, i+t[0], j+t[1])
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ret++
				fmt.Println(i, j)
				dfs(grid, i, j)
			}
		}
	}
	fmt.Println(grid[0])
	return ret
}
