package problemset

func buddyStrings(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal { //字符串相等需要有频次大于1的字符
		seen := [26]bool{}
		for _, ch := range s {
			if seen[ch-'a'] {
				return true
			}
			seen[ch-'a'] = true
		}
		return false
	}
	//只能交换一次，所以除了两个交换位置外，其他的字符都一样
	first, second := -1, -1
	for i := range s {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
