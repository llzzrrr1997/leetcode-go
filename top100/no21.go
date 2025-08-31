package top100

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
				cur = cur.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
				cur = cur.Next
			}
		} else if list1 != nil {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else if list2 != nil {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}
	return head.Next
}
