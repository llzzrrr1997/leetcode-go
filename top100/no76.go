package top100

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	sCnt, tCnt := [128]int{}, [128]int{}
	left := 0
	right := 0
	ret := ""
	for _, v := range t {
		tCnt[v]++
	}
	for right < len(s) {
		sCnt[s[right]]++
		for minWindowCover(sCnt, tCnt) {
			if len(ret) == 0 || len(ret) > right-left+1 {
				ret = s[left : right+1]
			}
			sCnt[s[left]]--
			left++
		}
		right++
	}
	return ret
}

func minWindowCover(s, t [128]int) bool {
	for i := 'a'; i <= 'z'; i++ {
		if s[i] < t[i] {
			return false
		}
	}

	for i := 'A'; i <= 'Z'; i++ {
		if s[i] < t[i] {
			return false
		}
	}
	return true
}
