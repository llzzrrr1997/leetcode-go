package top100

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		//nums[nums[i]-1] != nums[i]  如果2个交互的数相等不判断这个条件会死循环
		for nums[i] > 0 && nums[i] <= n && nums[i]-1 != i && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return n + 1
}
