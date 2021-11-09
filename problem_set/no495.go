package problemset

func findPoisonedDuration(timeSeries []int, duration int) int {
	ret := 0
	start := timeSeries[0]
	for i := 1; i < len(timeSeries); i++ {
		cur := start + duration
		if cur < timeSeries[i] {
			ret += duration
		} else {
			ret += timeSeries[i] - timeSeries[i-1]
		}
		start = timeSeries[i]
	}
	ret += duration
	return ret
}
