package foundation

//no334
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	small, mid := 2147483647, 2147483647
	for i := 0; i < n; i++ {
		if small >= nums[i] {
			small = nums[i]
		} else if mid >= nums[i] {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}

//238
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right, ret := make([]int, n), make([]int, n), make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		ret[i] = left[i] * right[i]
	}
	return ret
}

//no560
func subarraySum(nums []int, k int) int {
	cnt, pre := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		cnt += m[pre-k]
		m[pre] += 1
	}
	return cnt
}
