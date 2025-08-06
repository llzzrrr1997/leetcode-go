package problemset

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func removeElement2(nums []int, val int) int {
	i := 0
	for k := 0; k < len(nums); k++ {
		if nums[k] != val {
			nums[i] = nums[k]
			i++
		}
	}
	return i
}
