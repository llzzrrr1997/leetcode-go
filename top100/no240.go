package top100

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	m := len(matrix)
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
