package code_top

//leetcode 206
func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{Next: nil}
	for head != nil {
		tmp := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = tmp
	}
	return newHead.Next
}
