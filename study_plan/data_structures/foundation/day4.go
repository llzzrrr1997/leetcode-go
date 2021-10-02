package foundation

import "sort"

//no240
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	i, j := 0, col-1
	for i < row && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

//no435
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return n
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ret, right := 1, intervals[0][1]
	for i := 1; i < n; i++ {
		if right <= intervals[i][0] {
			ret++
			right = intervals[i][1]
		}
	}
	return n - ret
}
