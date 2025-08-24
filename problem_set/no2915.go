package problemset

import (
	"fmt"
	"math"
)

func lengthOfLongestSubsequence(nums []int, target int) int {

	m := make(map[string]int)
	var dfs func(i int, t int) int
	dfs = func(i int, t int) int {
		if t < 0 {
			return math.MinInt
		}
		if t == 0 {
			return 0
		}
		if i < 0 {
			return math.MinInt
		}
		k := fmt.Sprintf("%d-%d", i-1, t)
		v1, ok := m[k]
		if !ok {
			v1 = dfs(i-1, t)
			m[k] = v1
		}
		k = fmt.Sprintf("%d-%d", i-1, t-nums[i])
		v2, ok := m[k]
		if !ok {
			v2 = dfs(i-1, t-nums[i])
			m[k] = v2
		}

		return max(v1, v2+1)
	}

	ret := dfs(len(nums)-1, target)
	if ret < 0 {
		return -1
	}
	return ret

}

func lengthOfLongestSubsequence2(nums []int, target int) int {
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
