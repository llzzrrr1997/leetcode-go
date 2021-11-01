package problemset

func distributeCandies(candyType []int) int {
	n := len(candyType)
	set := make(map[int]bool)
	for _, val := range candyType {
		set[val] = true
	}
	if len(set) > n/2 {
		return n / 2
	}
	return len(set)
}
