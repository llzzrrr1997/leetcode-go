package problemset

func hIndex(citations []int) int {
	if len(citations) == 1 && citations[0] == 0 {
		return 0
	}
	right := len(citations) - 1
	i := 1
	for len(citations) >= i {
		if citations[right] < i {
			return i - 1
		}
		right--
		i++
	}
	return len(citations)
}
