package problemset

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
			dfs(j + 1) //如果选 nums[j] 添加到 path 末尾，那么下一个要添加到 path 末尾的数，就要在 nums[j+1] 到 nums[n−1] 中枚举了。
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return ret
}

//枚举子集（答案）的第一个数选谁，第二个数选谁，第三个数选谁，依此类推。
//
//dfs 中的 i 表示现在要枚举选 nums[i] 到 nums[n−1] 中的一个数，添加到 path 末尾。
//
//如果选 nums[j] 添加到 path 末尾，那么下一个要添加到 path 末尾的数，就要在 nums[j+1] 到 nums[n−1] 中枚举了。
