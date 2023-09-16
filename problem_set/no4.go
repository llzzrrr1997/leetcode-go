package problemset

import "sort"

// 直接合并2个数组然后进行排序，最后直接返回中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}
	return float64(nums[len(nums)/2])
}
