package problemset

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	ret := 0
	left := 0
	right := len(plants) - 1
	a := capacityA
	b := capacityB
	for left <= right {
		if left == right {
			if a >= b {
				if a < plants[left] {
					ret++
				}
			} else {
				if b < plants[left] {
					ret++
				}
			}
			left++
			break
		}
		if a >= plants[left] {
			a -= plants[left]
		} else {
			a = capacityA
			ret++
			a -= plants[left]
		}
		if b >= plants[right] {
			b -= plants[right]
		} else {
			b = capacityB
			ret++
			b -= plants[right]
		}
		left++
		right--
	}
	return ret
}
