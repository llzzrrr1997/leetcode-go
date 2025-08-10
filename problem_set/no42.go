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

func trap2(height []int) int {
	pre := make([]int, len(height))
	suf := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if i > 0 {
			pre[i] = max(pre[i-1], height[i])
		} else {
			pre[i] = height[i]
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		if i >= len(height)-1 {
			suf[i] = height[i]
		} else {
			suf[i] = max(height[i], suf[i+1])
		}
	}
	ret := 0
	for i := 0; i < len(height); i++ {
		h := min(suf[i], pre[i])
		if h > height[i] {
			ret += h - height[i]
		}
	}
	return ret
}
