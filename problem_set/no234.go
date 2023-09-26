package problemset

func isPalindrome234(head *ListNode) bool {
	//找出中点，进行反转最后判断
	//奇数节点则在中间节点一个进行反转
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	newHead := reverse234(slow)
	for newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}
	return true
}

func reverse234(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
