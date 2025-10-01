package top100

func jump(nums []int) int {
	cnt := 0
	maxLen := 0
	curMaxLen := 0
	for i, v := range nums[:len(nums)-1] {
		maxLen = max(maxLen, i+v)
		if curMaxLen == i {
			cnt++
			curMaxLen = maxLen
		}
	}

	return cnt
}
