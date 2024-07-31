package top100

func trap(height []int) int {
	n := len(height)
	preMax := make([]int, n)
	sufMax := make([]int, n)
	preMax[0] = height[0]
	sufMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		preMax[i] = height[i]
		if preMax[i] < preMax[i-1] {
			preMax[i] = preMax[i-1]
		}
	}
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = height[i]
		if sufMax[i] < sufMax[i+1] {
			sufMax[i] = sufMax[i+1]
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		min := preMax[i]
		if min > sufMax[i] {
			min = sufMax[i]
		}
		ret += min - height[i]
	}
	return ret
}

func trap2(height []int) int {
	n := len(height)
	ret := 0
	left := 0
	right := n - 1
	preMax := 0
	sufMax := 0
	for left <= right {
		if height[left] > preMax {
			preMax = height[left]
		}
		if height[right] > sufMax {
			sufMax = height[right]
		}
		if height[left] < height[right] {
			ret += preMax - height[left]
			left++
		} else {
			ret += sufMax - height[right]
			right--
		}
	}
	return ret
}
