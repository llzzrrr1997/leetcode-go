package problemset

// 差分数组工具类
type Difference struct {
	// 差分数组
	diff []int
}

func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

/* 给闭区间 [i, j] 增加 val（可以是负数）*/
func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d *Difference) Result() []int {
	res := make([]int, len(d.diff))
	// 根据差分数组构造原数组
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001) //题目最多的数量
	dif := NewDifference(nums)
	for _, trip := range trips {
		// 乘客数量
		val := trip[0]
		// 第 trip[1] 站乘客上车
		i := trip[1]
		// 第 trip[2] 站乘客已经下车，
		// 即乘客在车上的区间是 [trip[1], trip[2] - 1]
		j := trip[2] - 1
		// 进行区间操作
		dif.Increment(i, j, val)
	}
	ret := dif.Result()
	for _, val := range ret {
		if val > capacity {
			return false
		}
	}
	return true
}
