package problemset

// 双指针比较
func maxArea(height []int) int {
	start, end := 0, len(height)-1
	ret := 0
	for start < end {
		width := end - start
		low := height[start]
		if height[start] > height[end] {
			low = height[end]
			end--
		} else {
			start++
		}
		if width*low > ret {
			ret = width * low
		}
	}
	return ret
}
