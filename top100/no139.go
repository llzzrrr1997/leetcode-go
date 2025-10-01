package top100

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := range s {
		curVal := false
		for _, word := range wordDict {
			if len(word) <= i+1 && word == s[i+1-len(word):i+1] && !curVal {
				curVal = dp[i+1-len(word)]
			}
		}
		dp[i+1] = curVal
	}
	fmt.Println(dp)
	return dp[n]
}
