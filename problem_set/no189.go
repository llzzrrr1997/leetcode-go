package problemset

func rotate(nums []int, k int) {
	newNums := make([]int, 2*len(nums))
	copy(newNums, nums)
	copy(newNums[len(nums):], nums)
	k = k % len(nums)
	k = len(nums) - k
	for i := k; i < len(nums)+k; i++ {
		nums[i-k] = newNums[i]
	}
}
