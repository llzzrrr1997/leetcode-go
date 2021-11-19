package problemset

func findLHS(nums []int) int {
	cntMap := make(map[int]int)
	for _, val := range nums {
		cntMap[val]++
	}
	ret := 0
	for key, val := range cntMap {
		val2, ok := cntMap[key+1]
		if ok && ret < val2+val {
			ret = val2 + val
		}
	}
	return ret
}
