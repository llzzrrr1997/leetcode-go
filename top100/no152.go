package top100

func maxProduct(nums []int) int {
	n := len(nums)
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMin[0] = nums[0]
	dpMax[0] = nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		dpMax[i] = max(max(dpMax[i-1]*x, dpMin[i-1]*x), x)
		dpMin[i] = min(min(dpMin[i-1]*x, dpMax[i-1]*x), x)
	}
	ret := dpMax[0]
	for i := 1; i < n; i++ {
		ret = max(ret, dpMax[i])
	}
	return ret
}
