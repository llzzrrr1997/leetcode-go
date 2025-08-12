package problemset

import "math"

func minimumTime(time []int, totalTrips int) int64 {

	minTime := math.MaxInt
	for _, v := range time {
		minTime = min(minTime, v)
	}

	left := minTime - 1
	right := minTime * totalTrips
	for left < right {
		mid := (right + left) / 2
		sum := 0
		for _, v := range time {
			sum += mid / v
		}

		if sum >= totalTrips {
			right = mid
		} else {
			left = mid
		}

	}
	return int64(right)
}
