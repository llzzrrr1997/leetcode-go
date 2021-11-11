package problemset

import "math"

//切换英文才知道意思，题意太难懂了
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			min := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + max(dp[i][k-1], dp[k+1][j])
				if cost < min {
					min = cost
				}
			}
			dp[i][j] = min
		}
	}
	return dp[1][n]
}
