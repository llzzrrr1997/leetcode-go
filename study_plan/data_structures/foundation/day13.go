package foundation

//no24
func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	pre := newHead
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return newHead.Next
			}
		}
		next := tail.Next
		head, tail = reverseNodeList(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return newHead.Next
}

func reverseNodeList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}

//no143
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
