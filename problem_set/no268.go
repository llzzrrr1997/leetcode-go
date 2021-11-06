package problemset

import "sort"

func missingNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return nums[i] - 1
		}
	}
	return n
}

func missingNumber2(nums []int) int {
	n := len(nums)
	sum := 0
	for _, val := range nums {
		sum += val
	}
	total := (n + 1) * n / 2
	return total - sum
}
