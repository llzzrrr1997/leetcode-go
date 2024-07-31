package top100

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ret := 0
	for left < right {
		low := height[left]
		if low > height[right] {
			low = height[right]
			right--
		} else {
			left++
		}
		tmp := low * (right - left + 1)
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}
