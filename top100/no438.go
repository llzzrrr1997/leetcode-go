package top100

// 定长滑动窗口，窗口为p的字符
func findAnagrams(s string, p string) []int {
	window := [26]byte{}
	windowLen := len(p)
	for i := range p {
		window[p[i]-'a']++
	}
	ret := make([]int, 0)
	left, right, n := 0, 0, len(s)
	curWindow := [26]byte{}
	for right < n {
		curWindow[s[right]-'a']++
		if right-left+1 < windowLen {
			right++
			continue
		}
		if window == curWindow {
			ret = append(ret, left)
		}
		curWindow[s[left]-'a']--
		curWindow[s[right]-'a']-- //因为循环会设置s[right]-'a'++，所以执行到这里需要减去一次
		left++
	}
	return ret
}
