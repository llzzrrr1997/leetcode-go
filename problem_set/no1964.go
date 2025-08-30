package problemset

func longestObstacleCourseAtEachPosition(nums []int) []int {
	//最长递增子序列，常规dp超时
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
	return dp
}
