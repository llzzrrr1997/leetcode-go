package problemset

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	ret := 0
	for _, val := range arr {
		dp[val] = dp[val-difference] + 1
		if dp[val] > ret {
			ret = dp[val]
		}
	}
	return ret
}
