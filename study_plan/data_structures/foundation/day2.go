package foundation

import "sort"

//no75
func sortColors(nums []int) {
	cntZero, cntOne, cntTwo := 0, 0, 0
	for i := range nums {
		if nums[i] == 0 {
			cntZero++
		} else if nums[i] == 1 {
			cntOne++
		} else {
			cntTwo++
		}
	}
	i := 0
	for cntZero > 0 || cntOne > 0 || cntTwo > 0 {
		if cntZero > 0 {
			cntZero--
			nums[i] = 0
		} else if cntOne > 0 {
			cntOne--
			nums[i] = 1
		} else {
			cntTwo--
			nums[i] = 2
		}
		i++
	}
}

//no56
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ret = append(ret, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			if end < intervals[i][1] {
				end = intervals[i][1]
			}
		}
	}
	ret = append(ret, []int{start, end})
	return ret
}

//no706
type MyHashMap struct {
	data []int
	seen []bool
}

func Constructor() MyHashMap {
	return MyHashMap{
		data: make([]int, 1000001),
		seen: make([]bool, 1000001),
	}
}

func (m *MyHashMap) Put(key int, value int) {
	m.data[key] = value
	m.seen[key] = true
}

func (m *MyHashMap) Get(key int) int {
	if m.seen[key] {
		return m.data[key]
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	m.seen[key] = false
}
