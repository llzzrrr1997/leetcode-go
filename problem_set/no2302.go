package problemset

func countSubarrays2302(nums []int, k int64) int64 {
	left := 0
	ret := 0
	sum := 0
	for right := range nums {
		sum += nums[right]
		point := sum * (right - left + 1)
		for int64(point) >= k {
			sum -= nums[left]
			left++
			point = sum * (right - left + 1)
		}
		ret += right - left + 1 //左端点到右端点个子数组个数

	}
	return int64(ret)
}
