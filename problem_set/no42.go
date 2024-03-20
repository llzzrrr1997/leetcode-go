package problemset

func trap(height []int) int {
	//接雨水：按列求，然后累加
	//对于每一列，只需要看它左边和右边的最大列，然后取最小的值，如果当前列小于最小的值，则可以接雨水当前值-最小的值
	left := make([]int, len(height))
	right := make([]int, len(height))
	ans := 0
	leftMax := height[0]
	for i := 0; i < len(height); i++ {
		left[i] = max(height[i], leftMax)
		leftMax = left[i]
	}
	rightMax := height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		right[i] = max(height[i], rightMax)
		rightMax = right[i]
	}
	for i := 0; i < len(height); i++ {
		tmp := min(left[i], right[i])
		if height[i] < tmp {
			ans += tmp - height[i]
		}
	}
	return ans
}
