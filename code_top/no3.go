package code_top

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ret, start, end := 0, 0, 0
	window := make(map[uint8]int)
	for end < n {
		window[s[end]]++
		for window[s[end]] > 1 {
			window[s[start]]--
			start++
		}
		ret = max(ret, end-start+1)
		end++
	}
	return ret
}
