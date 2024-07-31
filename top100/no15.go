package top100

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums) - 1
	ret := make([][]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] { //过滤相同的数字
			continue
		}
		j := i + 1
		k := n
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[k-1] == nums[k] {
					k--
				}
				for j < k && nums[j+1] == nums[j] {
					j++
				}
				j++
				k--
			} else if tmp > 0 {
				k--
			} else {
				j++

			}
		}
	}
	return ret
}
