package problemset

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	visited := make([]bool, len(nums))
	cur := 0
	var helper func(nums []int, index int, cur *int) bool
	helper = func(nums []int, index int, cur *int) bool {
		if k == 0 {
			return true
		}
		if *cur == target {
			var newCur = 0
			k--
			ret := helper(nums, 0, &newCur)
			if !ret {
				k++
			}
			return ret
		}
		for i := index; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if *cur+nums[i] > target {
				continue
			}
			visited[i] = true
			*cur += nums[i]
			if helper(nums, i+1, cur) {
				return true
			}
			visited[i] = false
			*cur -= nums[i]
		}
		return false
	}
	return helper(nums, 0, &cur)
}
