package top100

import "sort"

func partitionLabels(s string) []int {
	rangeList := [][2]int{} //每个字母的区间
	minMap := make(map[int32]int)
	maxMap := make(map[int32]int)
	for i, v := range s {
		_, ok := minMap[v]
		if !ok {
			minMap[v] = i
		}
		maxMap[v] = i
	}
	for v, start := range minMap {
		end := maxMap[v]
		rangeList = append(rangeList, [2]int{start, end})
	}
	sort.Slice(rangeList, func(i, j int) bool {
		if rangeList[i][0] == rangeList[j][0] {
			return rangeList[i][1] < rangeList[j][1]
		}
		return rangeList[i][0] < rangeList[j][0]
	})
	//合并区间
	ret := make([]int, 0)
	start := rangeList[0][0]
	end := rangeList[0][1]
	for i := 1; i < len(rangeList); i++ {
		if rangeList[i][0] > end {
			ret = append(ret, end-start+1)
			start = rangeList[i][0]
			end = rangeList[i][1]
		}
		end = max(end, rangeList[i][1])
	}
	ret = append(ret, end-start+1)
	return ret
}
