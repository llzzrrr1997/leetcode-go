package top100

func isPalindrome(head *ListNode) bool {
	val := make([]int, 0)
	for head != nil {
		val = append(val, head.Val)
		head = head.Next
	}
	left, right := 0, len(val)-1
	for left < right {
		if val[left] != val[right] {
			return false
		}
		left++
		right--
	}
	return true
}
