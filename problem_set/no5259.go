package problemset

func calculateTax(brackets [][]int, income int) float64 {
	var ret float64
	pre := []int{0, 0}
	n := len(brackets)
	for i := 0; i < n; i++ {
		if income < brackets[i][0] {
			ret += float64((income-pre[0])*brackets[i][1]) / 100.0
			break
		} else {
			ret += float64((brackets[i][0]-pre[0])*brackets[i][1]) / 100.0
			pre = brackets[i]
		}
	}
	return ret
}
