package problemset

func getIntersectionNodeLCR023(headA, headB *ListNode) *ListNode {
	cntA, cntB := 0, 0
	curA, curB := headA, headB
	for curA != nil || curB != nil {
		if curA != nil {
			cntA++
			curA = curA.Next
		}
		if curB != nil {
			cntB++
			curB = curB.Next
		}
	}
	if cntA < cntB {
		cntA, cntB = cntB, cntA
		headA, headB = headB, headA
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
