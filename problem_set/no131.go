package problemset

func partition131(s string) [][]string {
	ret := [][]string{}
	cur := []string{}
	var dfs func(i, startI int)

	dfs = func(i, startI int) {
		if i == len(s) {
			newCur := make([]string, len(cur))
			copy(newCur, cur)
			ret = append(ret, newCur)
			return
		}

		if i < len(s)-1 {
			dfs(i+1, startI)
		}

		curS := s[startI : i+1]
		if check131(curS) {
			cur = append(cur, curS)
			dfs(i+1, i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, 0)
	return ret
}

func check131(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 枚举子串结束位置
func partition131_2(s string) [][]string {
	ret := [][]string{}
	cur := []string{}
	var dfs func(s string)

	dfs = func(ss string) {
		if len(ss) == 0 {
			newCur := make([]string, len(cur))
			copy(newCur, cur)
			ret = append(ret, newCur)
			return
		}
		for i := 0; i < len(ss); i++ {
			tmp := ss[:i+1]
			if check131(tmp) {
				cur = append(cur, tmp)
				dfs(ss[i+1:])
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs(s)
	return ret
}
