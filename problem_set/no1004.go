package problemset

func longestOnes(nums []int, k int) int {
	ret := 0
	left := 0
	zeroIndex := make(map[int]bool)
	for right := range nums {
		if nums[right] == 0 {
			zeroIndex[right] = true
		}
		t := len(zeroIndex)
		for t > k {
			if zeroIndex[left] {
				delete(zeroIndex, left)
				t--
			}
			left++
		}
		ret = max(ret, right-left+1)
	}
	return ret
}
