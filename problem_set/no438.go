package problemset

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	targetMap := [26]int32{}
	for _, val := range p {
		targetMap[val-'a']++
	}
	curMap := [26]int32{}
	for i := 0; i < len(p); i++ {
		curMap[s[i]-'a']++
	}
	ret := []int{}
	if curMap == targetMap {
		ret = append(ret, 0)
	}
	n := len(p)
	for i := n; i < len(s); i++ {
		curMap[s[i]-'a']++
		curMap[s[i-n]-'a']--
		if curMap == targetMap {
			ret = append(ret, i-n+1)
		}
	}
	return ret
}

func findAnagramsWindow(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	need := make(map[rune]int)   // 记录目标字符串 t 中每个字符的出现次数
	window := make(map[rune]int) // 记录窗口中每个字符出现的次数
	for _, c := range p {        // 初始化 need
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	res := []int{} // 记录结果
	for right < len(s) {
		c := rune(s[right])
		right++

		if _, ok := need[c]; ok { // 开始滑动窗口，进行窗口数据更新
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) { // 判断左侧窗口是否要收缩
			if valid == len(need) { // 如果当前窗口符合条件，把窗口左侧索引加入 res
				res = append(res, left)
			}
			d := rune(s[left])
			left++

			if _, ok := need[d]; ok { // 进行窗口内数据的更新
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
