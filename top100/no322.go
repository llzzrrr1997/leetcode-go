package top100

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		curVal := 100001
		for _, coin := range coins {
			if i >= coin && dp[i-coin] < curVal {
				curVal = dp[i-coin] + 1
			}
		}
		dp[i] = curVal
	}
	if dp[amount] == 100001 {
		return -1
	}
	return dp[amount]
}
