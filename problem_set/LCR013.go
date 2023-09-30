package problemset

type NumMatrixLCR013 struct {
	// 定义：preSum[i][j] 记录 matrix 中子矩阵 [0, 0, i-1, j-1] 的元素和
	preSum [][]int
}

func ConstructorLCR013(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix)+1)
	for i := 0; i < len(preSum); i++ {
		preSum[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{preSum: preSum}
}

func (this *NumMatrixLCR013) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}
