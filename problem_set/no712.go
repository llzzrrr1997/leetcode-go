package problemset

func minimumDeleteSum(s1 string, s2 string) int {
	n := len(s1)
	m := len(s2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for j := 0; j <= m; j++ {
		if j > 0 {
			dp[0][j] = dp[0][j-1] + int(s2[j-1])
		}

	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] //字符一样不删
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1])) //不一样时要么删除i（dp[i-1][j]）,要么删除j（dp[i][j-1]）
			}
		}
	}
	return dp[n][m]
}
