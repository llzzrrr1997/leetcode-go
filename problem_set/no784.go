package problemset

func letterCasePermutation(s string) []string {
	ret := []string{}
	cur := []byte{}
	ss := []byte(s)
	var dfs func(i int) //处理s[i:]的数据
	dfs = func(i int) {
		if len(cur) == len(ss) {
			ret = append(ret, string(cur))
			return
		}
		for j := i; j < len(ss); j++ {
			if ss[j] >= 'a' && ss[j] <= 'z' {
				cur = append(cur, ss[j])
				dfs(j + 1)
				cur = cur[:len(cur)-1]

				cur = append(cur, ss[j]-'a'+'A')
				dfs(j + 1)
				cur = cur[:len(cur)-1]

			} else if ss[j] >= 'A' && ss[j] <= 'Z' {
				cur = append(cur, ss[j])
				dfs(j + 1)
				cur = cur[:len(cur)-1]

				cur = append(cur, ss[j]-'A'+'a')
				dfs(j + 1)
				cur = cur[:len(cur)-1]
			} else {
				cur = append(cur, ss[j])
				dfs(j + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs(0)

	return ret
}

func letterCasePermutation2(s string) []string {
	ret := []string{}
	ss := []byte(s)
	var dfs func(i int) //处理s[i:]的数据
	dfs = func(i int) {
		ret = append(ret, string(ss))
		if i == len(ss) {
			return
		}
		for j := i; j < len(ss); j++ {
			if ss[j] >= 'a' && ss[j] <= 'z' {
				ss[j] = ss[j] - 'a' + 'A'
				dfs(j + 1)
				ss[j] = ss[j] + 'a' - 'A'
			} else if ss[j] >= 'A' && ss[j] <= 'Z' {
				ss[j] = ss[j] - 'A' + 'a'
				dfs(j + 1)
				ss[j] = ss[j] + 'A' - 'a'
			}
		}
	}
	return ret
}
