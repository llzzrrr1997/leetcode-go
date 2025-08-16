package problemset

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//1->2<-3<-4<-5
//   |
//   V
//   nil

func reverseList2(head *ListNode) *ListNode {
	newHead := &ListNode{}
	for head != nil {
		tmp := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = tmp
	}
	return newHead.Next
}
