package problemset

func longestSemiRepetitiveSubstring(s string) int {
	ret := 0
	left := 0
	same := 0
	for right := range s {
		if right > 0 && s[right] == s[right-1] {
			same++
		}
		for same > 1 && right > left {
			if s[left] == s[left+1] {
				same--
			}
			left++
		}
		ret = max(ret, right-left+1)
	}
	return ret
}
