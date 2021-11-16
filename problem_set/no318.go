package problemset

func maxProduct(words []string) int {
	ret := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if maxProductHelper(words[i], words[j]) {
				cur := len(words[i]) * len(words[j])
				if ret < cur {
					ret = cur
				}
			}
		}
	}
	return ret
}
func maxProductHelper(s1, s2 string) bool {
	m := make(map[int32]bool)
	for _, val := range s1 {
		m[val] = true
	}
	for _, val := range s2 {
		if m[val] {
			return false
		}
	}
	return true
}
