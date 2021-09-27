package foundation

//no119
func getRow(rowIndex int) []int {
	ans := make([][]int, rowIndex+1)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans[rowIndex]
}

//no48
func rotate(matrix [][]int) {
	//水平翻转
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	//对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

//no59
func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	left, right, top, bootom := 0, n-1, 0, n-1
	for num := 1; num <= n*n; {
		//从左往右
		for i := left; i <= right; i++ {
			ret[top][i] = num
			num++
		}
		top++
		//从上往下
		for i := top; i <= bootom; i++ {
			ret[i][right] = num
			num++
		}
		right--
		//从右往左
		for i := right; i >= left; i-- {
			ret[bootom][i] = num
			num++
		}
		bootom--
		//从下往上
		for i := bootom; i >= top; i-- {
			ret[i][left] = num
			num++
		}
		left++
	}
	return ret
}
