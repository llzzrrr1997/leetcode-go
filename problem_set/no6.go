package problemset

// 用 2 个变量保存方向，当垂直输出的行数达到了规定的目标行数以后，需要从下往上转折到第一行，循环中控制好方向
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	bytes := make([][]byte, numRows)
	for i := range bytes {
		bytes[i] = make([]byte, 0, numRows)
	}
	col := 0
	row := 0
	down := true
	for i := 0; i < len(s); i++ {
		bytes[row] = append(bytes[row], s[i])
		if row+1 == numRows {
			down = false
		}
		if row == 0 {
			down = true
		}
		if !down {
			col++
		}
		if down {
			row++
		} else {
			row--
		}
	}
	newS := make([]byte, 0, len(s))
	for i := 0; i < len(bytes); i++ {
		newS = append(newS, bytes[i]...)
	}
	return string(newS)
}
