package problemset

import "math"

func canPartition(nums []int) bool {
	s := 0
	for _, v := range nums {
		s += v
	}
	if s%2 != 0 {
		return false
	}

	return !(helper416(nums, s/2) == -1)

}

// 0-1背包，看2915
func helper416(nums []int, target int) int {
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}

	//dp[i][t] = max(dp[i-1][t-nums[i]]+1,dp[i-1][t])
	for i := 0; i <= len(nums); i++ {
		for t := 1; t <= target; t++ {
			if i-1 < 0 {
				dp[i][t] = math.MinInt
				continue
			}
			if t-nums[i-1] < 0 {
				dp[i][t] = dp[i-1][t]
				continue
			}
			dp[i][t] = max(dp[i-1][t-nums[i-1]]+1, dp[i-1][t])
		}
	}
	if dp[len(nums)][target] < 0 {
		return -1
	}

	return dp[len(nums)][target]
}
