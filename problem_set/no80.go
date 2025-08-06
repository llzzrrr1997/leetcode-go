package problemset

func removeDuplicates3(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0
	m := make(map[int]int, len(nums))
	m[nums[0]]++
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			m[nums[j]]++
		} else if m[nums[j]] < 2 {
			i++
			nums[i] = nums[j]
			m[nums[j]]++
		}
	}
	return i + 1
}

// 题解，用栈进行处理
func removeDuplicates4(nums []int) int {
	stackSize := 2
	for j := 2; j < len(nums); j++ {
		if nums[stackSize-2] != nums[j] {
			nums[stackSize] = nums[j]
			stackSize++
		}
	}
	return min(stackSize, len(nums))
}
