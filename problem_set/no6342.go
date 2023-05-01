package problemset

func firstCompleteIndex(arr []int, mat [][]int) int {
	type position struct {
		x, y int
	}
	row := len(mat)
	col := len(mat[0])
	xPMap := make([]int, row)
	yPMap := make([]int, col)
	pMap := make(map[int]position)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			pMap[mat[i][j]] = position{i, j}
		}
	}
	for i, v := range arr {
		xPMap[pMap[v].x]++
		if xPMap[pMap[v].x] == col {
			return i
		}
		yPMap[pMap[v].y]++
		if yPMap[pMap[v].y] == row {
			return i
		}
	}
	return 0
}

//4,3,5
//1,2,6

//1,4,5,2,6,3
