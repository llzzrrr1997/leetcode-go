package problemset

func twoSumLCR006(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		tmp := numbers[left] + numbers[right]
		if tmp == target {
			return []int{left, right}
		}
		if tmp > target {
			right--
		}
		if tmp < target {
			left++
		}
	}
	return []int{}
}
