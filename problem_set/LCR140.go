package problemset

func trainingPlan(head *ListNode, cnt int) *ListNode {
	cur := head
	for i := 0; i < cnt; i++ {
		cur = cur.Next
	}
	slow := head
	for cur != nil {
		slow = slow.Next
		cur = cur.Next
	}
	return slow
}
