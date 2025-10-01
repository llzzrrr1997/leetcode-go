package top100

import "math"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	ret := math.MaxInt
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			ret = min(ret, nums[mid])
			left = mid + 1
		} else {
			ret = min(ret, nums[mid])
			right = mid - 1
		}
	}
	return ret
}
