package top100

import "math"

func maxSubArray(nums []int) int {
	ret := math.MinInt
	preSum := 0    //当前前缀和
	minPreSum := 0 //最小前缀和
	for _, v := range nums {
		preSum += v
		tmp := preSum - minPreSum
		if tmp > ret {
			ret = tmp
		}
		if minPreSum > preSum {
			minPreSum = preSum
		}
	}
	return ret
}
