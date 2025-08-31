package top100

func swapPairs(head *ListNode) *ListNode {
	n := 0
	k := 2
	newHead := &ListNode{Next: head}
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	var p0, pre *ListNode = newHead, nil
	cur = newHead.Next
	//k个一组
	for i := n; i >= k; i = i - k {
		for j := 0; j < k; j++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

		next := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = next
	}

	return newHead.Next
}
