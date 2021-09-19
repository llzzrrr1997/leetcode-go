package getting_started

type ListNode struct {
	Val  int
	Next *ListNode
}

//no141
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//no21
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				cur.Next = l2
				cur = cur.Next
				l2 = l2.Next
			} else {
				cur.Next = l1
				cur = cur.Next
				l1 = l1.Next
			}
			continue
		}
		if l1 != nil {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
		if l2 != nil {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	return newHead.Next
}

//no203
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	cur := newHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return newHead.Next
}
