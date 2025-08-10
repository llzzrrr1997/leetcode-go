package problemset

import "sort"

// 18. 四数之和 https://leetcode.cn/problems/4sum/solution/ji-zhi-you-hua-ji-yu-san-shu-zhi-he-de-z-1f0b/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] { //重复的数字
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] { //重复的数字（i+1）
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				s := nums[i] + nums[j] + nums[left] + nums[right]
				if s == target {
					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right { //重复的数字
						if nums[left] == nums[left-1] {
							left++
						} else {
							break
						}
					}
					right--
					for left < right { //重复的数字
						if nums[right] == nums[right+1] {
							right--
						} else {
							break
						}
					}
				} else if s > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return ret
}
