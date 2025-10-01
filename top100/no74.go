package top100

import "sort"

func searchMatrixNo74(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if !(matrix[i][len(matrix[i])-1] >= target && matrix[i][0] <= target) {
			continue
		}
		j := sort.SearchInts(matrix[i], target)
		if matrix[i][j] == target {
			return true
		} else {
			return false
		}
	}
	return false
}
