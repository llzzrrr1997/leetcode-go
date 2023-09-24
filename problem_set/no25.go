package problemset

// k个一组反转，可以反转前k个，然后k+1节点再反转k个
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	k1Node := head
	for i := 0; i < k; i++ {
		//k1Node = k1Node.Next //注意可能出现反转整个链表情况所以不能在这赋值，比如 1->2 k=2
		if k1Node == nil {
			return head
		}
		k1Node = k1Node.Next //注意可能出现反转整个链表情况，比如 1->2 k=2
	}
	newHead := reverseN(head, k)
	head.Next = reverseKGroup(k1Node, k)
	return newHead
}
