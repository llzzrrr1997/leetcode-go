package problemset

func countSubarrays(nums []int, k int) int64 {
	m := make(map[int]int)
	maxNum := 0
	for _, v := range nums {
		maxNum = max(v, maxNum)
	}
	ret := 0
	left := 0
	for right := range nums {
		m[nums[right]]++
		for m[maxNum] >= k {
			ret += len(nums) - right
			m[nums[left]]--
			left++
		}
	}
	return int64(ret)
}
