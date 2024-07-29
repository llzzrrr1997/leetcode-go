package top100

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	ret := 0
	for v := range m {
		delete(m, v)
		tmpRet := 1
		left := v
		for {
			if m[left-1] {
				tmpRet++
				left = left - 1
				delete(m, left)
			} else {
				break
			}
		}
		right := v
		for {
			if m[right+1] {
				tmpRet++
				right = right + 1
				delete(m, right)
			} else {
				break
			}
		}
		if ret < tmpRet {
			ret = tmpRet
		}
	}
	return ret
}
