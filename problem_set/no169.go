package problemset

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
	ret := 0
	cnt := 0
	for _, val := range nums {
		if cnt == 0 {
			cnt++
			ret = val
		} else if ret == val {
			cnt++
		} else if ret != val {
			cnt--
		}
	}
	return ret
}
