package top100

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{Next: nil}
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = newHead.Next
		newHead.Next = tmp
	}
	return newHead.Next
}
