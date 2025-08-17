package problemset

func permute(nums []int) [][]int {
	ret := [][]int{}
	v := make(map[int]bool)

	cur := []int{}
	var dfs func(i int)
	dfs = func(i int) { //已经保存了多少个数字
		if i == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ret = append(ret, tmp)
			return
		}
		for j := 0; j < len(nums); j++ {
			if !v[nums[j]] {
				cur = append(cur, nums[j])
				v[nums[j]] = true
				dfs(i + 1)
				cur = cur[:len(cur)-1]
				v[nums[j]] = false
			}
		}

	}

	dfs(0)
	return ret
}
