package top100

import (
	"sort"
	"strconv"
	"strings"
)

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	cur := []int{}
	retMap := make(map[string]bool)
	var dfs func(t int)
	dfs = func(t int) {
		if t == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			sort.Ints(tmp)
			k := strings.Builder{}
			for _, curV := range tmp {
				k.WriteString(strconv.Itoa(curV))
			}
			if retMap[k.String()] {
				return
			}
			retMap[k.String()] = true
			ret = append(ret, tmp)
		}
		if t <= 0 {
			return
		}
		for _, val := range candidates {
			cur = append(cur, val)
			dfs(t - val)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(target)
	return ret
}
