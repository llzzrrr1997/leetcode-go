package problemset

func numRollsToTarget(n int, k int, target int) int {
	//dp[i][j] 代表 扔 i 个骰子和为 j
	//dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j - 2] + ... + dp[i - 1][j - k]
	//边界条件： dp[1][k] = 1 ( 1<= k <= min(target, k) )
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	min := k
	if k > target {
		min = target
	}
	//初始化base
	for i := 1; i <= min; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= target; j++ {
			for m := 1; j-m >= 1 && m <= k; m++ {
				dp[i][j] = (dp[i-1][j-m] + dp[i][j]) % 1000000007
			}
		}
	}
	return dp[n][target]
}
