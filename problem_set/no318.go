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

func maxProduct2(words []string) int {
	wordCharMap := make([][26]int, 0, len(words))
	for _, val := range words {
		cur := [26]int{}
		for _, b := range val {
			cur[b-'a'] = 1
		}
		wordCharMap = append(wordCharMap, cur)
	}
	ret := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			hasCommon := false
			for k := 0; k < 26; k++ {
				if wordCharMap[i][k] == wordCharMap[j][k] && wordCharMap[i][k] == 1 {
					hasCommon = true
					break
				}
			}
			if !hasCommon && len(words[i])*len(words[j]) > ret {
				ret = len(words[i]) * len(words[j])
			}
		}
	}
	return ret
}
