package top100

func sortColors(nums []int) {
	cntZero, cntOne, cntTwo := 0, 0, 0
	for i := range nums {
		if nums[i] == 0 {
			cntZero++
		} else if nums[i] == 1 {
			cntOne++
		} else {
			cntTwo++
		}
	}
	i := 0
	for cntZero > 0 || cntOne > 0 || cntTwo > 0 {
		if cntZero > 0 {
			cntZero--
			nums[i] = 0
		} else if cntOne > 0 {
			cntOne--
			nums[i] = 1
		} else {
			cntTwo--
			nums[i] = 2
		}
		i++
	}
}
