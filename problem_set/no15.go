package problemset

import (
	"fmt"
	"sort"
)

// O(n2)超时
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	difMap := make(map[int]int)
	for i, val := range nums {
		difMap[val] = i
	}
	set := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			index, ok := difMap[0-nums[i]-nums[j]]
			if ok && index != i && index != j {
				tmp := []int{nums[i], nums[j], 0 - nums[i] - nums[j]}
				sort.Ints(tmp)
				key := fmt.Sprintf("%d%d%d", tmp[0], tmp[1], tmp[2])
				if !set[key] {
					ret = append(ret, tmp)
					set[key] = true
				}

			}
		}
	}
	return ret
}

// 排序后用二分查找
func threeSum2(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { //去除重复的
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			tmp := nums[i] + nums[start] + nums[end]
			if tmp == 0 {
				ret = append(ret, []int{nums[i], nums[start], nums[end]})
				for start < end && nums[start] == nums[start+1] { //去除重复的
					start++
				}
				for end > start && nums[end] == nums[end-1] { //去除重复的
					end--
				}
				start++
				end--
			} else if tmp < 0 {
				start++
			} else if tmp > 0 {
				end--
			}
		}
	}
	return ret
}
