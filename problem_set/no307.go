package problemset

import "fmt"

type NumArray307 struct {
	data   []int
	preSum []int
}

func Constructor307(nums []int) NumArray307 {
	a := NumArray307{}
	a.data = nums
	preSum := make([]int, len(nums)+1)
	sum := 0
	for i, val := range nums {
		sum += val
		preSum[i+1] = sum
	}
	a.preSum = preSum
	fmt.Println(preSum)
	return a
}

func (this *NumArray307) Update(index int, val int) {
	dif := val - this.data[index]
	this.data[index] = val
	for i := index; i <= len(this.data); i++ {
		this.preSum[i] += dif
	}
	fmt.Println(this.preSum)
}

func (this *NumArray307) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
