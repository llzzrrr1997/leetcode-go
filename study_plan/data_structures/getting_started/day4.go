package getting_started

//no566
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}
	ret := make([][]int, r)
	p, q := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if len(ret[p]) == 0 {
				ret[p] = make([]int, c)
			}
			ret[p][q] = mat[i][j]
			q++
			if q == c {
				q = 0
				p++
			}
		}
	}
	return ret
}

//no118
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
