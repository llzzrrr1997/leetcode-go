package problemset

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ret := 0
	prod := 1
	left := 0
	for right, x := range nums {
		prod = prod * x
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ret += right - left + 1 //left ... right中子数字个数都是满足乘积小于k
	}
	return ret
}
