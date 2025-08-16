package problemset

func findMin(nums []int) int {
	left, right := 0, len(nums)-1 // 闭区间 [0, n-1]
	// 循环条件：闭区间不为空
	for left < right {
		mid := left + (right-left)/2
		// 比较中间值与右边界，判断最小值在左侧还是右侧
		if nums[mid] < nums[right] {
			// 最小值在左侧（包含mid），收缩右边界
			right = mid
		} else {
			// 最小值在右侧（不包含mid），收缩左边界
			left = mid + 1
		}
	}
	// 循环结束时，left == right，即为最小值索引
	return nums[left]
}
