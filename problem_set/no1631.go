package problemset

import "math"

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	// 定义：从 (0, 0) 到 (i, j) 的最小体力消耗是 effortTo[i][j]
	effortTo := make([][]int, m)
	for i := range effortTo {
		effortTo[i] = make([]int, n)
		for j := range effortTo[i] {
			effortTo[i][j] = math.MaxInt32
		}
	}
	effortTo[0][0] = 0
	q := make([][]int, 0)
	q = append(q, []int{0, 0, 0})
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		curX, curY, curEffortFromStart := cur[0], cur[1], cur[2]
		if curEffortFromStart > effortTo[curX][curY] {
			continue
		}
		for _, neighbor := range adj1631(heights, curX, curY) {
			nextX, nextY := neighbor[0], neighbor[1]
			effortToNextNode := max1631(
				effortTo[curX][curY],
				abs1631(heights[curX][curY]-heights[nextX][nextY]),
			)
			if effortToNextNode < effortTo[nextX][nextY] {
				effortTo[nextX][nextY] = effortToNextNode
				q = append(q, []int{nextX, nextY, effortToNextNode})
			}
		}
	}
	return effortTo[m-1][n-1]
}

func adj1631(heights [][]int, x, y int) [][]int {
	m, n := len(heights), len(heights[0])
	res := [][]int{}
	if x > 0 {
		res = append(res, []int{x - 1, y})
	}
	if y > 0 {
		res = append(res, []int{x, y - 1})
	}
	if x < m-1 {
		res = append(res, []int{x + 1, y})
	}
	if y < n-1 {
		res = append(res, []int{x, y + 1})
	}
	return res
}

func abs1631(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max1631(a, b int) int {
	if a > b {
		return a
	}
	return b
}
