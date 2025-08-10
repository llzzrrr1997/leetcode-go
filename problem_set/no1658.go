package problemset

// 逆向思维+滑动窗口
func minOperations(nums []int, x int) int {
	sum := 0
	minNums := nums[0]
	for _, v := range nums {
		sum += v
		minNums = min(v, minNums)
	}
	if sum < x || minNums > x {
		return -1
	}
	dif := sum - x
	maxLen := -1
	left := 0
	sum = 0
	for right := range nums {
		sum += nums[right]
		for sum > dif {
			sum -= nums[left]
			left++
		}
		if sum == dif {
			maxLen = max(right-left+1, maxLen)
		}
	}
	if maxLen < 0 {
		return -1
	}
	return len(nums) - maxLen
}
