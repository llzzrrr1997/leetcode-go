package top100

import "sort"

func searchRange(nums []int, target int) []int {
	start := sort.SearchInts(nums, target)
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := sort.SearchInts(nums, target+1)
	return []int{start, end - 1}

}
