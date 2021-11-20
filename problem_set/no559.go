package problemset

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	ret := 0
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) > 0 {
		ret++
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			for _, val := range cur.Children {
				q = append(q, val)
			}
		}
	}
	return ret
}
