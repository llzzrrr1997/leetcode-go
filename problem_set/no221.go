package problemset

//["1","0","1","0","0"],
//["1","0","1","1","1"],
//["1","1","1","1","1"],
//["1","0","0","1","0"]

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)) //右下角正方形的边长
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				//dp[i][j] = min(dp[i-1][j],dp[i-1][j-1],dp[i][j-1])+1
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + 1
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			maxSide = max(dp[i][j], maxSide)
		}
	}

	return maxSide * maxSide
}
