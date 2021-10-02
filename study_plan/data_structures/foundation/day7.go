package foundation

import (
	"fmt"
	"sort"
	"strings"
)

//no290
func wordPattern(pattern string, s string) bool {
	list := strings.Split(s, " ")
	c2s := make(map[byte]string)
	s2c := make(map[string]byte)
	if len(list) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		s := list[i]
		if (c2s[pattern[i]] != "" && c2s[pattern[i]] != s) || (s2c[s] != 0 && s2c[s] != pattern[i]) {
			return false
		}
		c2s[pattern[i]] = s
		s2c[s] = pattern[i]
	}
	return true
}

//no763
func partitionLabels(s string) []int {
	firstM, lastM := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		_, ok := firstM[s[i]]
		if !ok {
			firstM[s[i]], lastM[s[i]] = i, i
		}
		if ok {
			lastM[s[i]] = i
		}
	}
	peer := make([][]int, 0)
	for k, val := range firstM {
		peer = append(peer, []int{val, lastM[k]})
	}
	fmt.Println(peer)
	sort.Slice(peer, func(i, j int) bool {
		return peer[i][0] < peer[j][0]
	})
	ret := make([]int, 0)
	left, right := peer[0][0], peer[0][1]
	for i := 1; i < len(peer); i++ {
		if peer[i][0] > right {
			ret = append(ret, right-left+1)
			left = peer[i][0]
			right = peer[i][1]
			continue
		}
		if right < peer[i][1] {
			right = peer[i][1]
		}
	}
	ret = append(ret, right-left+1)
	return ret
}
