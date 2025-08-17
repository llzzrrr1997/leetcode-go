package problemset

func letterCombinations(digits string) []string {
	ret := []string{}
	var mapping = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if digits == "" {
		return ret
	}

	var dfs func(path []byte)
	dfs = func(path []byte) {
		if len(path) == len(digits) {
			ret = append(ret, string(path))
			return
		}
		for _, c := range mapping[digits[len(path)]-'0'] {
			path = append(path, byte(c))
			dfs(path)
			path = path[:len(path)-1]
		}
	}

	dfs([]byte{})
	return ret
}
