package problemset

import "strings"

type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{
		make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	ret := 0
	for k, val := range this.m {
		if strings.HasPrefix(k, prefix) {
			ret += val
		}
	}
	return ret
}
