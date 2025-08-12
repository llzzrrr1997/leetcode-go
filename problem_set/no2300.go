package problemset

import "sort"

// 超时
func successfulPairs(spells []int, potions []int, success int64) []int {
	ret := make([]int, 0, len(spells))
	m := make(map[int]int)
	for _, v := range spells {
		tmpRet, ok := m[v]
		if ok {
			ret = append(ret, tmpRet)
			continue
		}
		for _, vv := range potions {
			if v*vv > int(success) {
				tmpRet++
			}
		}
		m[v] = tmpRet
		ret = append(ret, tmpRet)
	}
	return ret
}

func successfulPairs2(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ret := make([]int, 0, len(spells))
	for _, v := range spells {
		minPotion := (int(success)-1)/v + 1
		i := sort.SearchInts(potions, minPotion)
		ret = append(ret, len(potions)-i)
	}
	return ret
}
