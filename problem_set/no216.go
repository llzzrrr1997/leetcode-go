package problemset

func combinationSum3(k int, sum int) [][]int {
	ret := [][]int{}
	cur := []int{}
	n := 9
	var dfs func(i int)

	dfs = func(i int) {
		if len(cur) == k {
			//和为sum则添加
			tmp := 0
			for _, v := range cur {
				tmp += v
			}
			if tmp == sum {
				newCur := make([]int, len(cur))
				copy(newCur, cur)
				ret = append(ret, newCur)
			}

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
