package top100

import "fmt"

func productExceptSelf(nums []int) []int {
	//1,2,3,4
	pre := make([]int, len(nums))  //1,1,2,6
	suff := make([]int, len(nums)) //24,12,4,1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			pre[i] = 1
		} else {
			pre[i] = pre[i-1] * nums[i-1]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			suff[i] = 1
		} else {
			suff[i] = suff[i+1] * nums[i+1]
		}
	}
	fmt.Println(pre, suff)
	ret := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		ret = append(ret, pre[i]*suff[i])
	}
	return ret
}
