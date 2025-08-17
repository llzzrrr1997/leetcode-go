package problemset

func solveNQueens(n int) [][]string {
	ret := [][]string{}
	cur := make([][]rune, n)
	for i := 0; i < n; i++ {
		cur[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			cur[i][j] = '.'
		}
	}
	var dfs func(i, j int, qN int)

	dfs = func(i, j int, qN int) {
		if qN == n {
			newCur := make([]string, n)
			for i := 0; i < n; i++ {
				newCur[i] = string(cur[i])
			}
			ret = append(ret, newCur)
			return
		}

		for m := i; m < n; m++ {
			for k := j; k < n; k++ {
				if check51(m, k, cur) {
					cur[m][k] = 'Q'
					dfs(m+1, 0, qN+1)
					cur[m][k] = '.'
				}
			}
		}

	}

	dfs(0, 0, 0)
	return ret
}

func check51(i, j int, cur [][]rune) bool {

	//横不能出现皇后
	for k := j - 1; k >= 0; k-- {
		if cur[i][k] == 'Q' {
			return false
		}
	}

	//竖不能出现皇后
	for k := i - 1; k >= 0; k-- {
		if cur[k][j] == 'Q' {
			return false
		}
	}

	//左上斜不能出现皇后
	for m, n := i-1, j-1; m >= 0 && n >= 0; {
		if cur[m][n] == 'Q' {
			return false
		}
		m--
		n--
	}

	//右上斜不能出现皇后
	for m, n := i-1, j+1; m >= 0 && n < len(cur); {
		if cur[m][n] == 'Q' {
			return false
		}
		m--
		n++
	}
	return true
}
