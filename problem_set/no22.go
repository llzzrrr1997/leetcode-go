package problemset

func generateParenthesis(n int) []string {
	ret := []string{}
	cur := []rune{}

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if right == n && left == n {
			if check22(string(cur)) {
				ret = append(ret, string(cur))
			}
			return
		}
		if left < n {
			cur = append(cur, '(')
			dfs(left+1, right)
			cur = cur[:len(cur)-1]
		}
		if right < n {
			cur = append(cur, ')')
			dfs(left, right+1)
			cur = cur[:len(cur)-1]
		}

	}

	dfs(0, 0)

	return ret
}

func check22(s string) bool {
	cur := []rune{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cur = append(cur, '(')
		} else {
			if len(cur) == 0 {
				return false
			}
			cur = cur[:len(cur)-1]
		}
	}
	return true
}
