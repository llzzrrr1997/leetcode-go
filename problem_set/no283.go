package problemset

func moveZeroes(nums []int) {
	index := moveZeroesHelper(nums)
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
	return
}

func moveZeroesHelper(nums []int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
