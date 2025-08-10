package problemset

func maxSubarrayLength(nums []int, k int) int {
	m := make(map[int]int)
	ret := 0
	left := 0
	for right, x := range nums {
		m[x]++
		for m[x] > k {
			m[nums[left]]--
			left++
		}
		ret = max(ret, right-left+1)
	}
	return ret
}
