package top100

func setZeroes(matrix [][]int) {
	row := make(map[int]bool)
	col := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := range row {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}
	for j := range col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][j] = 0
		}
	}
}
