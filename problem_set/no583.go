package problemset

func minDistance583(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] //字符一样不删
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1 //不一样时要么删除i（dp[i-1][j]）,要么删除j（dp[i][j-1]）
			}
		}
	}
	return dp[n][m]
}
