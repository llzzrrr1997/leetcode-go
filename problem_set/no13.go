package problemset

// 按照题目的描述，可以总结如下规则：
//
// 罗马数字由 I,V,X,L,C,D,M 构成；
// 当小值在大值的左边，则减小值，如 IV=5-1=4；
// 当小值在大值的右边，则加小值，如 VI=5+1=6；
// 由上可知，右值永远为正，因此最后一位必然为正。
// 一言蔽之，把一个小值放在大值的左边，就是做减法，否则为加法。
func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	total, lastNum, n := 0, 0, len(s)
	for i := 0; i < n; i++ {
		curStr := s[n-(i+1) : n-i]
		curNum := roman[curStr]
		if curNum < lastNum {
			total -= curNum
		} else {
			total += curNum
		}
		lastNum = curNum
	}
	return total
}
