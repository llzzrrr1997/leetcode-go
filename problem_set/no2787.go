package problemset

import (
	"math"
)

// dp[i][n] = dp[i-1][n]+dp[i-1][n-i^x]
func numberOfWays(n int, x int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	mod := 1000000007
	dp[0][0] = 1
	for i := 0; i <= n; i++ {
		for t := 0; t <= n; t++ {
			p := int(math.Pow(float64(i), float64(x)))
			if i < 1 {
				continue
			}
			if t < p {
				dp[i][t] = dp[i-1][t] % mod
				continue
			}
			dp[i][t] = (dp[i-1][t] % mod) + (dp[i-1][t-p] % mod)
		}
	}

	return dp[n][n]
}
