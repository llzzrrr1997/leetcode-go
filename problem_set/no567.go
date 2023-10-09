package problemset

func checkInclusion(s2 string, s1 string) bool {
	if len(s1) < len(s2) {
		return false
	}
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range s2 {
		need[s2[i]]++
	}
	left, right := 0, 0
	count := 0
	for right < len(s1) {
		c := s1[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				count++
			}
		}
		for right-left >= len(s2) {
			if count == len(need) {
				return true
			}
			c = s1[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					count--
				}
				window[c]--
			}
		}
	}
	return false
}
