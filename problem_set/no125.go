package problemset

import "strings"

func isPalindrome125(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for !((s[left] >= 'a' && s[left] <= 'z') || (s[left] >= '0' && s[left] <= '9')) && left < right {
			left++
		}
		for !((s[right] >= 'a' && s[right] <= 'z') || (s[right] >= '0' && s[right] <= '9')) && left < right {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
