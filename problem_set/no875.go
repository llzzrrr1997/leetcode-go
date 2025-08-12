package problemset

func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	left := 0  //最小值
	right := 0 //最大值
	for _, v := range piles {
		right = max(v, right)
	}

	for left+1 < right {
		mid := (left + right) / 2
		sum := n
		for _, p := range piles {
			sum += (p - 1) / mid
		}
		if sum <= h {
			right = mid // 循环不变量：恒为 true
		} else {
			left = mid // 循环不变量：恒为 false
		}
	}
	return right
}
