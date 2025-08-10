package problemset

import "sort"

// 统计和小于目标的下标对数目 https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/solution/onlogn-pai-xu-shuang-zhi-zhen-by-endless-qk40/
func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	ret := 0
	for left < right {
		tmp := nums[left] + nums[right]
		if tmp < target {
			ret += right - left //num[left] + num[right---left+1] 都是小于target
			left++
		} else {
			right--
		}
	}
	return ret
}
