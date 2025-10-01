package top100

func orangesRotting(grid [][]int) int {
	ret := 0
	d := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	q := [][2]int{}
	fresh := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	for fresh > 0 && len(q) > 0 {
		tmpQ := [][2]int{}
		ret++
		for _, v := range q {
			for _, dir := range d {
				i := v[0] + dir[0]
				j := v[1] + dir[1]
				if i < 0 || j < 0 || j >= len(grid[0]) || i >= len(grid) {
					continue
				}
				if grid[i][j] == 1 {
					grid[i][j] = 2
					fresh--
					tmpQ = append(tmpQ, [2]int{i, j})
				}
			}
		}
		q = tmpQ
	}

	if fresh > 0 {
		return -1
	}

	return ret
}
