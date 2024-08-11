package top100

func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] += preSum[i-1] + nums[i-1]
	}
	ret := 0
	cntMap := make(map[int]int)
	for _, s := range preSum {
		ret += cntMap[s-k] //s-k + s = k
		cntMap[s]++
	}
	return ret
}
