package problemset

func maximumCount(nums []int) int {
	neg := lowerBound(nums, 0) //负数的数量
	pos := len(nums) - lowerBound(nums, 1)
	return max(neg, pos)
}

// 返回最小满足nums[i] >= target 的i
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
