package problemset

// 优先选择大的数字，解题思路采用贪心算法。将 1-3999 范围内的罗马数字从大到小放在数组中，从头选择到尾，即可把整数转成罗马数字。
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ret := ""
	i := 0
	for num > 0 {
		for values[i] > num {
			i++
		}
		ret += symbols[i]
		num -= values[i]
	}
	return ret
}
