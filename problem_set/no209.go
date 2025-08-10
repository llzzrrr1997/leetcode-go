package problemset

import "math"

func minSubArrayLen(target int, nums []int) int {
	ret := math.MaxInt
	s := 0
	left := 0
	for i, x := range nums {
		s += x
		for s >= target {
			ret = min(ret, i-left+1)
			s -= nums[left]
			left++
		}
	}
	if ret == math.MaxInt {
		return 0
	}
	return ret
}
