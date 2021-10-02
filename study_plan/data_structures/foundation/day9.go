package foundation

//no187
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n <= 10 {
		return []string{}
	}
	cntMap := make(map[string]int)
	for i := 10; i <= n; i++ {
		cntMap[s[i-10:i]]++
	}
	ret := make([]string, 0)
	for k, val := range cntMap {
		if val > 1 {
			ret = append(ret, k)
		}
	}
	return ret
}

//no5
func longestPalindromeNo5(s string) string {
	//中心扩展
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left, right := expandAroundCenter(s, i, i)
		if right-left > end-start {
			start, end = left, right
		}
		left, right = expandAroundCenter(s, i, i+1)
		if right-left > end-start {
			start, end = left, right
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
