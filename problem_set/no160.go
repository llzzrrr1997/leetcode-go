package problemset

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//统计链表的长度，然后比较差值n提前先遍历n次最后一起遍历出交点
	curA, curB := headA, headB
	cntA, cntB := 1, 1
	for curA.Next != nil || curB.Next != nil {
		if curA.Next != nil {
			curA = curA.Next
			cntA++
		}
		if curB.Next != nil {
			curB = curB.Next
			cntB++
		}
	}
	if curA != curB { //2个末尾节点不一致肯定没有相交
		return nil
	}
	if cntB > cntA {
		headB, headA = headA, headB
		cntB, cntA = cntA, cntB
	}
	dif := cntA - cntB
	for i := 0; i < dif; i++ {
		headA = headA.Next
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
