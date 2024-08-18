package top100

func rotate48(matrix [][]int) {
	tmp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		tmp[i] = make([]int, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			tmp[j][len(matrix)-i-1] = matrix[i][j]
		}
	}
	copy(matrix, tmp)
}
