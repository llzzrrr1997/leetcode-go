package top100

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cntA, cntB := 1, 1
	newHeadA := headA
	newHeadB := headB
	for newHeadA.Next != nil || newHeadB.Next != nil {
		if newHeadA.Next != nil {
			cntA++
			newHeadA = newHeadA.Next
		}
		if newHeadB.Next != nil {
			cntB++
			newHeadB = newHeadB.Next
		}
	}
	if newHeadA != newHeadB {
		return nil
	}
	dif := cntA - cntB
	if cntA < cntB {
		dif = cntB - cntA
		headA, headB = headB, headA
	}
	for i := 0; i < dif; i++ {
		headA = headA.Next
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
