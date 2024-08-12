package top100

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for _, v := range intervals {
		if v[0] > end {
			ret = append(ret, []int{start, end})
			start, end = v[0], v[1]
		} else if v[1] > end {
			end = v[1]
		}
	}
	ret = append(ret, []int{start, end})
	return ret
}
