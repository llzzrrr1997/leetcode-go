package problemset

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	even := false
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
		if even { //反转tmpVals
			left, right := 0, len(tmpVals)-1
			for left < right {
				tmpVals[left], tmpVals[right] = tmpVals[right], tmpVals[left]
				left++
				right--
			}
		}
		ret = append(ret, tmpVals)
		even = !even
	}

	return ret
}
