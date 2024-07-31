package top100

func moveZeroes(nums []int) {
	i := 0
	for k := range nums {
		if nums[k] != 0 {
			nums[i] = nums[k]
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
