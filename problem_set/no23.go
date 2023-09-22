package problemset

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	newList := make([]*ListNode, 0, len(lists)/2)
	left, right := 0, len(lists)-1
	if len(lists)%2 == 1 {
		newList = append(newList, lists[0])
		left = 1
	}
	for left < right {
		tmp := mergeTwoLists(lists[left], lists[right])
		newList = append(newList, tmp)
		left++
		right--
	}
	return mergeKLists(newList)
}
