package problemset

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			} else if j+1 == len(triangle[i]) {
				dp[i][j] = triangle[i][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = triangle[i][j] + min(dp[i-1][j], dp[i-1][j-1])
			}

		}
	}
	ret := dp[n-1][0]
	for i := 1; i < len(triangle[n-1]); i++ {
		ret = min(ret, dp[n-1][i])
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
