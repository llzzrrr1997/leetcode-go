package code_top

import "sort"

//可以使用快排或者堆排进行排序后返回数据
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
