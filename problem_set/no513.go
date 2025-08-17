package problemset

func findBottomLeftValue(root *TreeNode) int {
	//层序遍历得到最后一层的第一个数据
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

	return ret[len(ret)-1][0]
}

// 从右节点先入队，最后层序遍历的最后值就是结果
func findBottomLeftValue2(root *TreeNode) int {
	q := make([]*TreeNode, 0, 1)
	node := root
	q = append(q, root)
	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}

	return node.Val
}
