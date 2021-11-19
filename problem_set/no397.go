package problemset

func integerReplacement(n int) int {
	m := make(map[int]int)
	var dfs func(n int) int
	dfs = func(n int) int {
		if n == 1 {
			return 0
		}
		val, ok := m[n]
		if ok {
			return val
		}
		ret := 0
		if n % 2 == 0 {
			ret =  dfs(n/2)
		}else {
			addRet := dfs(n+1)
			difRet := dfs(n-1)
			if addRet > difRet {
				ret = difRet
			}else {
				ret = addRet
			}
		}
		m[n] = ret+1
		return ret+1
	}
	return dfs(n)
}