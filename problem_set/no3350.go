package problemset

func maxIncreasingSubarrays(nums []int) int {
	//答案一：连续递增子数组的长度/2
	//答案二：上一个连续递增子数组和当前连续递增子数组的长度中的最小值
	//取答案一和答案二的最大值
	ret := 0
	preCnt := 0
	curCnt := 0
	for i := 0; i < len(nums); i++ {
		curCnt++
		if i == len(nums)-1 || nums[i] >= nums[i+1] {
			ret1 := curCnt / 2
			ret2 := min(preCnt, curCnt)
			ret = max(max(ret1, ret2), ret)
			preCnt = curCnt
			curCnt = 0
		}
	}

	return ret
}
