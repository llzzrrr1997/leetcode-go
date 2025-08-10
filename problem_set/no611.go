package problemset

import (
	"sort"
)

// 611. 有效三角形的个数 https://leetcode.cn/problems/valid-triangle-number/solution/zhuan-huan-cheng-abcyong-xiang-xiang-shu-1ex3/
func triangleNumber(nums []int) int {
	ret := 0
	sort.Ints(nums)
	n := len(nums)
	for i := 2; i < n; i++ { //固定最大值
		j := 0
		k := i - 1
		for j < k {
			if nums[j]+nums[k] > nums[i] {
				ret += k - j //j+1 到 k都是大于 nums[i]
				k--
			} else {
				j++
			}
		}
	}
	return ret
}

// a > b - c
func triangleNumber1(nums []int) int {
	ret := 0
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] == 0 {
			continue
		}
		j := i + 1
		for k := j + 1; k < n; k++ {
			for nums[k]-nums[j] >= nums[i] {
				j++
			}
			ret += k - j
		}
	}
	return ret
}
