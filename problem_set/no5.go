package problemset

// 动态
// 动态规划。定义 dp[i][j] 表示从字符串第 i 个字符到第 j 个字符这一段子串是否是回文串。由回文串的性质可以得知，回文串去掉一头一尾相同的字符以后，剩下的还是回文串。
// 所以状态转移方程是 dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])。
// 注意特殊的情况，j - i == 1 的时候，即只有 2 个字符的情况，只需要判断这 2 个字符是否相同即可。j - i == 2 的时候，即只有 3 个字符的情况，只需要判断除去中心以外对称的 2 个字符是否相等。每次循环动态维护保存最长回文串即可。时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- { //从后开始遍历确保dp[i+1][j-1]已经初始化
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

// 中心扩散，从当前字符往外判断是否字符一样直到不一样为止，注意 奇数长度的回文则以当前字符为中间，偶数长度的回文则当前字符和下一个字符为中间
func longestPalindrome2(s string) string {
	res := ""
	f := func(j, k int) {
		cur := ""
		for j >= 0 && k <= len(s)-1 && s[j] == s[k] {
			cur = s[j : k+1]
			j--
			k++
		}
		if len(cur) > len(res) {
			res = cur
		}
	}
	for i := 0; i < len(s); i++ {
		f(i, i)
		f(i, i+1)
	}
	return res
}
