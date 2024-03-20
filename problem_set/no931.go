package problemset

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(matrix[i]))
		if i == 0 {
			for j := 0; j < n; j++ {
				dp[i][j] = matrix[i][j]
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j+1])
			} else if j+1 == len(matrix[i]) {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j-1], dp[i-1][j])
			} else {
				dp[i][j] = matrix[i][j] + min(min(dp[i-1][j], dp[i-1][j-1]), dp[i-1][j+1])
			}
		}
	}
	ret := dp[n-1][0]
	for i := 1; i < len(matrix[n-1]); i++ {
		ret = min(ret, dp[n-1][i])
	}
	return ret
}
