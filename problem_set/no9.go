package problemset

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// 不使用字符串的解法
func isPalindrome2(x int) bool {
	//小于0肯定不是回文数字
	if x < 0 {
		return false
	}
	//如果以0结尾，只有0是回文
	if x != 0 && x%10 == 0 {
		return false
	}
	//回文一半的数字，然后与剩下的数字进行对比
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	//如果是偶数的长度则两个值要相等，如果是奇数则回文一半的数revertedNumber/10 要相等
	return x == revertedNumber || x == revertedNumber/10
}
