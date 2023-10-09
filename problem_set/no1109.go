package problemset

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	df := NewDifference(nums)
	for _, val := range bookings {
		df.Increment(val[0]-1, val[1]-1, val[2])
	}
	return df.Result()
}
