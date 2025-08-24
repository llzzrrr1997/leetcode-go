package problemset

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1) //dp[i][j]为选coins[i]构造成j amount的数量;
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= amount; j++ {
			if coins[i] > j {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j] + dp[i+1][j-coins[i]]
			}
		}
	}
	return dp[n][amount]
}
