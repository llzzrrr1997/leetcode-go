package foundation

type ListNode struct {
	Val  int
	Next *ListNode
}

//no2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	m := 0
	for l1 != nil || l2 != nil || m != 0 {
		if l1 != nil {
			m += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			m += l2.Val
			l2 = l2.Next
		}
		cur := &ListNode{
			Val:  m % 10,
			Next: nil,
		}
		m = m / 10
		if head == nil {
			head = cur
		}
		if tail == nil {
			tail = cur
		} else {
			tail.Next = cur
			tail = tail.Next
		}
	}
	return head
}

//no142
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//先判断存在环
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
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
