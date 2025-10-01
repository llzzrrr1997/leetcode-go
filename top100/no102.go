package top100

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	q := make([]*TreeNode, 0, 1)
	if root != nil {
		q = append(q, root)
	}
	for len(q) != 0 {
		tmp := make([]*TreeNode, 0, 2*len(q))
		tmpInt := make([]int, 0, len(q))
		for len(q) != 0 {
			tmpInt = append(tmpInt, q[0].Val)
			if q[0].Left != nil {
				tmp = append(tmp, q[0].Left)
			}
			if q[0].Right != nil {
				tmp = append(tmp, q[0].Right)
			}
			q = q[1:]
		}
		q = tmp
		ret = append(ret, tmpInt)
	}

	return ret
}
