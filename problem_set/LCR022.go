package problemset

func detectCycleLCR022(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	//假设相遇的节点距离起点为m，slow走k个节点则fast走2k
	//fast比slow多走的k个都是在环里循环
	//slow到环起始节点为k-m
	//fast在环里循环的k-m次也是起始节点
	//所以相遇节点就是slow和fast重新从头走k-m步
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
