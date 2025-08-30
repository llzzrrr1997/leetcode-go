package problemset

func minimumOperations(nums []int) int {
	//最长递增子序列
	//n - 最长递增子序列就是答案
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ret := dp[0]
	for i := 0; i < n; i++ {
		ret = max(ret, dp[i])
	}
	return n - ret
}
