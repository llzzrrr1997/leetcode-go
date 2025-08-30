package problemset

// 最长回文子串等于 s和s的翻转后字符串的最长公共子序列
func longestPalindromeSubseq(s string) int {
	d := []byte(s)
	l, r := 0, len(d)-1
	for l < r {
		d[l], d[r] = d[r], d[l]
		l++
		r--
	}
	s1 := s
	s2 := string(d)
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = max(max(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][n]
}

func longestPalindromeSubseq2(s string) int {
	n := len(s)
	dp := make([][]int, n) //从i到j到最长回文子串的长度 dp[i][j] = max(dp[i+1][j],dp[i][j-1]) || dp[i+1][j-1]+2
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	//因为dp[i][j] 从i+1号j-1得到，所以i从n-1开始，j从0开始
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1 //一个的时候为1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}
