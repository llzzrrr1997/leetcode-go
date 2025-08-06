package problemset

// 超时
func maxProfit(prices []int) int {
	p := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			p = max(p, prices[j]-prices[i])
		}
	}
	return p
}

// 在i天卖出时，最大值就是i天的价格减去i之前的最小价格
func maxProfit2(prices []int) int {
	p := 0
	minP := prices[0]
	for i := 1; i < len(prices); i++ {
		if minP > prices[i] {
			minP = prices[i]
		} else {
			p = max(p, prices[i]-minP)
		}
	}
	return p
}
