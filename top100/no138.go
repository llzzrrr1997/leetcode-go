package top100

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		tmp := &Node{Val: cur.Val}
		m[cur] = tmp
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		newCur := m[cur]
		newCur.Next = m[cur.Next]
		newCur.Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]

}
