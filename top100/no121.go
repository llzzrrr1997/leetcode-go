package top100

func maxProfit(prices []int) int {
	ret := 0
	minP := prices[0]
	for _, v := range prices {
		ret = max(ret, v-minP)
		if minP > v {
			minP = v
		}
	}

	return ret
}
