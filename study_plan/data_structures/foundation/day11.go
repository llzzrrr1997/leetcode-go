package foundation

//no160
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	//先判断是否存在相交
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
	if curA != curB {
		return nil
	}
	curA, curB = headA, headB
	if cntA < cntB {
		curA, curB = curB, curA
		cntA, cntB = cntB, cntA
	}
	dif := cntA - cntB
	for i := 0; i < dif; i++ {
		curA = curA.Next
	}
	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}
	return curA
}

//no82
func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	cur := newHead
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tmp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return newHead.Next
}
