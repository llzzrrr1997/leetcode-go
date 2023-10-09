package problemset

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	count := 0
	left, right := 0, 0
	ret := ""
	for right < len(s) {
		c := s[right]
		window[c]++
		right++
		if window[c] == need[c] {
			count++
		}
		for len(need) == count {
			if ret == "" {
				ret = s[left:right]
			} else if right-left < len(ret) {
				ret = s[left:right]
			}
			leftC := s[left]
			left++
			window[leftC]--
			if window[leftC] < need[leftC] {
				count--
			}
		}
	}
	return ret
}

// 滑动窗口算法框架
//func slidingWindow(s string) {
//	// 用合适的数据结构记录窗口中的数据
//	window := make(map[byte]int)
//
//	left, right := 0, 0
//	for right < len(s) {
//		// c 是将移入窗口的字符
//		c := s[right]
//		window[c]++
//		// 增大窗口
//		right++
//		// 进行窗口内数据的一系列更新
//
//		/*** debug 输出的位置 ***/
//		// 注意在最终的解法代码中不要输出
//		// 因为 IO 操作很耗时，可能导致超时
//		fmt.Printf("window: [%d, %d)\n", left, right)
//		/********************/
//
//		// 判断左侧窗口是否要收缩
//		for left < right && window needs shrink {
//			// d 是将移出窗口的字符
//			d := s[left]
//			window[d]--
//			// 缩小窗口
//			left++
//			// 进行窗口内数据的一系列更新
//		}
//	}
//}
