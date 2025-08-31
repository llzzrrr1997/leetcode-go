package top100

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	fast := newHead
	slow := newHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
