package top100

func canJump(nums []int) bool {
	maxLen := 0
	for i, v := range nums {
		if i > maxLen {
			return false
		}
		maxLen = max(i+v, maxLen)
	}
	return true
}
