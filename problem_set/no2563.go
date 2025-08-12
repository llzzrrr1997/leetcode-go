package problemset

import (
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ret := 0
	for i, v := range nums {
		if v > lower {
			break
		}
		left := lower - v
		right := upper - v + 1
		leftL := sort.SearchInts(nums[:i], left)
		leftR := sort.SearchInts(nums[:i], right)
		if leftR < leftL {
			continue
		}
		ret += leftR - leftL
	}
	return int64(ret)
}

func countFairPairs2(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ret := 0
	for i, v := range nums {
		left := lower - v
		right := upper - v + 1
		leftL := sort.SearchInts(nums[i+1:], left)
		leftR := sort.SearchInts(nums[i+1:], right)
		if leftR < leftL {
			continue
		}
		// fmt.Println(leftL, leftR)
		ret += leftR - leftL
	}
	return int64(ret)
}
