package top100

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	newNums := make([]int, n)
	for i := 0; i < n; i++ {
		newNums[(i+k)%n] = nums[i]
	}
	copy(nums, newNums)
}
