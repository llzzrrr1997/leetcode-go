package problemset

import "math/rand"

type Solution struct {
	data   []int
	origin []int
}

func Constructor384(nums []int) Solution {
	s := Solution{
		data:   nums,
		origin: append([]int(nil), nums...),
	}
	return s
}

func (this *Solution) Reset() []int {
	this.data = append([]int(nil), this.data...)
	return this.origin
}

func (this *Solution) Shuffle() []int {
	sData := make([]int, len(this.data))
	for i := range sData {
		j := rand.Intn(len(this.data))
		sData[i] = this.data[j]
		this.data = append(this.data[:j], this.data[j+1:]...)
	}
	this.data = sData
	return sData
}
