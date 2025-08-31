package problemset

func gcdOfOddEvenSums(n int) int {
	oddSum := 0
	evenSum := 0
	for i := 0; i < 2*n; i = i + 2 {
		oddSum += i + 1
		evenSum += i + 2
	}
	for evenSum != 0 {
		tmp := evenSum
		evenSum = oddSum % evenSum
		oddSum = tmp
	}
	return oddSum
}

func partitionArray(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	cntM := make(map[int]int)
	for _, v := range nums {
		cntM[v]++
	}

	n := len(nums) / k

	for _, cnt := range cntM {
		if cnt > n {
			return false
		}
	}
	return true
}
