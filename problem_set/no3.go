package problemset

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ret, start, end := 0, 0, 0
	set := make(map[byte]int)
	hasDoubled := false
	for end < n {
		set[s[end]]++
		hasDoubled = checkDouble(set)
		if hasDoubled {
			for start <= end {
				set[s[start]]--
				start++
				hasDoubled = checkDouble(set)
				if !hasDoubled {
					break
				}
			}
		}
		cur := end - start + 1
		if cur > ret {
			ret = cur
		}
		end++
	}
	fmt.Println(ret)
	return ret
}

func checkDouble(m map[byte]int) bool {
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring2(s string) int {
	n := len(s)
	ret, start, end := 0, 0, 0
	set := make(map[byte]int)
	for end < n {
		set[s[end]]++
		for set[s[end]] > 1 { //滑动窗口内，只有最后end元素才会大于1
			set[s[start]]--
			start++
		}
		cur := end - start + 1
		if cur > ret {
			ret = cur
		}
		end++
	}
	return ret
}

func lengthOfLongestSubstring3(s string) int {
	n := len(s)
	ret, start, end := 0, 0, 0
	set := make(map[byte]int)
	for end < n {
		if set[s[end]] == 0 {
			set[s[end]]++
			cur := end - start + 1
			if cur > ret {
				ret = cur
			}
			end++
		} else {
			set[s[start]]--
			start++
		}

	}
	return ret
}
