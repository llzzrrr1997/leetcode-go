package top100

func exist(board [][]byte, word string) bool {
	ret := false
	v := make([][]bool, len(board))
	for i := range v {
		v[i] = make([]bool, len(board[0]))
	}
	var dfs func(i, j int, index int)
	dfs = func(i, j int, index int) {
		if ret {
			return
		}
		if index == len(word) {
			ret = true
			return
		}
		if i >= len(board) || i < 0 {
			return
		}
		if j >= len(board[0]) || j < 0 {
			return
		}

		if board[i][j] == word[index] && !v[i][j] { //选
			v[i][j] = true
			dfs(i, j-1, index+1)
			dfs(i, j+1, index+1)
			dfs(i-1, j, index+1)
			dfs(i+1, j, index+1)
			v[i][j] = false
			return
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, 0)
		}
	}

	return ret
}
