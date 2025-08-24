package problemset

import "fmt"

func deleteAndEarn(nums []int) int {
	m := make(map[int]int)
	minNum := nums[0]
	maxNum := nums[0]
	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			m[v] = 0
		}
		m[v]++
		minNum = min(minNum, v)
		maxNum = max(maxNum, v)
	}
	newNums := make([]int, 0, maxNum-minNum+1)
	for i := minNum; i <= maxNum; i++ {
		cnt, ok := m[i]
		if ok {
			newNums = append(newNums, i*cnt)
		} else {
			newNums = append(newNums, 0)
		}

	}
	fmt.Println(newNums)
	if len(newNums) == 1 {
		return newNums[0]
	}
	if len(newNums) == 2 {
		return max(newNums[0], newNums[1])
	}

	dp := make([]int, len(newNums))
	dp[0] = newNums[0]
	dp[1] = max(newNums[0], newNums[1])
	for i := 2; i < len(newNums); i++ { //打家劫舍的做法，nums为i*cnt的数组
		dp[i] = max(dp[i-1], dp[i-2]+newNums[i])
	}
	return dp[len(newNums)-1]

}
