package problemset

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	ret := make([][]int, 0)
	for len(q) > 0 {
		tmpVals := make([]int, 0)
		nextQ := make([]*TreeNode, 0)
		for len(q) > 0 {
			tmpVals = append(tmpVals, q[0].Val)
			if q[0].Left != nil {
				nextQ = append(nextQ, q[0].Left)
			}
			if q[0].Right != nil {
				nextQ = append(nextQ, q[0].Right)
			}
			q = q[1:]
		}
		q = nextQ
		ret = append(ret, tmpVals)
	}

	return ret
}
