package problemset

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
