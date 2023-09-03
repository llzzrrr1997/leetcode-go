package problemset

func mySqrt(x int) int {
	//二分查找x
	l, r := 0, x
	ret := 0
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
