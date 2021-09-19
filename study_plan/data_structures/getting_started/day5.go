package getting_started

//no36
func isValidSudoku(board [][]byte) bool {
	rowMap, colMap, boxMap := make([][]int, 9), make([][]int, 9), make([][]int, 9)
	for i := 0; i < 9; i++ {
		rowMap[i] = make([]int, 9)
		colMap[i] = make([]int, 9)
		boxMap[i] = make([]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				val := int(board[i][j] - '1')
				boxVal := (i/3)*3 + j/3
				if rowMap[i][val] == 0 && colMap[j][val] == 0 && boxMap[boxVal][val] == 0 {
					rowMap[i][val] = 1
					colMap[j][val] = 1
					boxMap[boxVal][val] = 1
				} else {
					return false
				}
			}
		}
	}
	return true
}

//no73
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 && !seen[i][j] {
				for p := 0; p < n; p++ {
					if matrix[i][p] != 0 {
						matrix[i][p] = 0
						seen[i][p] = true
					}
				}
				for p := 0; p < m; p++ {
					if matrix[p][j] != 0 {
						matrix[p][j] = 0
						seen[p][j] = true
					}
				}
				seen[i][j] = true
			}
		}
	}
	return
}
