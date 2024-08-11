package top100

func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	q := []int{}
	for i, num := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-q[0] >= k {
			q = q[1:]
		}
		if i >= k-1 {
			ret = append(ret, nums[q[0]])
		}
	}
	return ret
}
