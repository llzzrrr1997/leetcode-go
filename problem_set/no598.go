package problemset

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) < 1 {
		return m * n
	}
	minM := ops[0][0]
	minN := ops[0][1]
	for i := 1; i < len(ops); i++ {
		if minM > ops[i][0] {
			minM = ops[i][0]
		}
		if minN > ops[i][1] {
			minN = ops[i][1]
		}
	}
	return minN * minM
}
