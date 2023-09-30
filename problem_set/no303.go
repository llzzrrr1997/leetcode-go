package problemset

type NumArray struct {
	preSum []int
}

func ConstructorNO303(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{
		preSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
