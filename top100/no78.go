package top100

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	cur := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		newCur := make([]int, len(cur))
		copy(newCur, cur)
		ret = append(ret, newCur)
		for j := i; j < len(nums); j++ {
			cur = append(cur, nums[j])
			dfs(j + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return ret
}
