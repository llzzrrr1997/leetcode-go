package code_top

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ret, start, end := 0, 0, 0
	indexMap := make(map[byte]int) //下标map
	for ; end < n; end++ {
		index, ok := indexMap[s[end]]
		if ok && start < index {
			start = index
		}
		cur := end - start + 1
		ret = max(cur, ret)
		indexMap[s[end]] = end + 1
	}
	return ret
}
