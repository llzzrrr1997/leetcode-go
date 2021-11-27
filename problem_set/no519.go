package problemset

import (
	"math/rand"
)

type No519Solution struct {
	data  map[int]int
	total int
	m, n  int
}

func No519Constructor(m int, n int) No519Solution {
	s := No519Solution{}
	s.m, s.n = m, n
	s.total = m * n
	s.data = make(map[int]int)
	return s
}

func (this *No519Solution) Flip() []int {
	x := rand.Intn(this.total)
	this.total--
	index, ok := this.data[x]
	if !ok {
		index = x
	}
	total, ok := this.data[this.total]
	if !ok {
		total = this.total
	}
	this.data[x] = total
	return []int{index / this.n, index % this.n}
}

func (this *No519Solution) Reset() {
	this.total = this.m * this.n
	this.data = make(map[int]int)
}
