package problemset

func combine(n int, k int) [][]int {
	ret := [][]int{}
	cur := []int{}
	var dfs func(i int)

	dfs = func(i int) {
		if len(cur) == k {
			newCur := make([]int, len(cur))
			copy(newCur, cur)
			ret = append(ret, newCur)
			return
		}

		for j := i; j <= n; j++ {
			cur = append(cur, j)
			dfs(j + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(1)

	return ret
}
