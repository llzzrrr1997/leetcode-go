package getting_started

//no350
func intersect(nums1 []int, nums2 []int) []int {
	cnt1 := make(map[int]int)
	cnt2 := make(map[int]int)
	for i := range nums1 {
		cnt1[nums1[i]]++
	}
	for i := range nums2 {
		cnt2[nums2[i]]++
	}
	ret := make([]int, 0)
	for k, val := range cnt1 {
		val2, ok := cnt2[k]
		cnt := 0
		if ok {
			if val > val2 {
				cnt = val2
			} else {
				cnt = val
			}
		}
		for i := 0; i < cnt; i++ {
			ret = append(ret, k)
		}
	}
	return ret
}

//no121
func maxProfit(prices []int) int {
	ret := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		cur := prices[i] - min
		if cur > ret {
			ret = cur
		}
	}
	return ret
}
