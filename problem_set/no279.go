package problemset

import "math"

// dp[n]:数字n的最少完全平方数的数量
// dp[i] = 1(i为完全平方数)；min(dp[i-k]+dp[k])
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sqrt := int(math.Sqrt(float64(i)))
		if sqrt*sqrt == i {
			dp[i] = 1
			continue
		}
		dp[i] = math.MaxInt
		for j := 1; j*j < i; j++ { //k优化为平方
			dp[i] = min(dp[i], dp[i-j*j]+dp[j*j])
		}
	}
	return dp[n]
}
