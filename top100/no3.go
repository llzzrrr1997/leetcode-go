package top100

// 不定长滑动窗口
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left, right := 0, 0
	ret := 0
	for right < len(s) {
		m[s[right]]++
		if m[s[right]] > 1 {
			for left < right && m[s[right]] > 1 {
				m[s[left]]--
				left++
			}
		}
		tmp := right - left + 1
		if ret < tmp {
			ret = tmp
		}
		right++
	}
	return ret
}
