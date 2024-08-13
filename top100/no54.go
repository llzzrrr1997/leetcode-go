package top100

func spiralOrder(matrix [][]int) []int {
	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1
	count := len(matrix) * len(matrix[0]) //统计总数
	ret := make([]int, 0, count)
	for count > 0 {
		//从左到右
		for i := left; i <= right && count > 0; i++ {
			ret = append(ret, matrix[top][i])
			count--
		}
		top++
		//从上到下
		for i := top; i <= bottom && count > 0; i++ {
			ret = append(ret, matrix[i][right])
			count--
		}
		right--
		//从右到左
		for i := right; i >= left && count > 0; i-- {
			ret = append(ret, matrix[bottom][i])
			count--
		}
		bottom--
		//从下到上
		for i := bottom; i >= top && count > 0; i-- {
			ret = append(ret, matrix[i][left])
			count--
		}
		left++
	}
	return ret
}
