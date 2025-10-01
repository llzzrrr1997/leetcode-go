package top100

func canPartition(nums []int) bool {
	s := 0
	for _, v := range nums {
		s += v
	}
	if s%2 != 0 {
		return false
	}
	s = s / 2
	n := len(nums)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	for i, x := range nums {
		for j := 0; j <= s; j++ {
			dp[i+1][j] = j >= x && dp[i][j-x] || dp[i][j]
		}
	}

	return dp[n][s]
}
