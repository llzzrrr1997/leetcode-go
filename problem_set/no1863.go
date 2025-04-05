package problemset

func subsetXORSum(nums []int) int {
	return subsetXORSumHelper(nums, 0, 0)
}

func subsetXORSumHelper(nums []int, index, val int) int {
	if index == len(nums) {
		return val
	}
	return subsetXORSumHelper(nums, index+1, nums[index]^val) + subsetXORSumHelper(nums, index+1, val)
}
