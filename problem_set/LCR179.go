package problemset

func twoSumLCR179(price []int, target int) []int {
	left, right := 0, len(price)-1
	for left < right {
		tmp := price[left] + price[right]
		if tmp == target {
			return []int{price[left], price[right]}
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
