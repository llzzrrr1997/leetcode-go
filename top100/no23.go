package top100

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	left, right := 0, len(lists)-1
	newList := make([]*ListNode, 0, len(lists)/2+1)
	if len(lists)%2 == 1 {
		left = 1
		newList = append(newList, lists[0])
	}
	for left < right {
		tmp := mergeKListsHelper(lists[left], lists[right])
		newList = append(newList, tmp)
		left++
		right--
	}
	return mergeKLists(newList)
}

func mergeKListsHelper(list1, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return newHead.Next
}
