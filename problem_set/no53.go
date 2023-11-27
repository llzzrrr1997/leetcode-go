package problemset

import "math"

func maxSubArray(nums []int) int {
	//前缀和最大的差值
	ret := math.MinInt
	minPreSum := 0
	preSum := 0
	for _, v := range nums {
		preSum += v
		if ret < preSum-minPreSum {
			ret = preSum - minPreSum
		}
		if minPreSum > preSum {
			minPreSum = preSum
		}
	}
	return ret
}
