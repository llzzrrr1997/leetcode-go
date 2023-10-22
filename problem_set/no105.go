package problemset

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(pre []int, preStart int, preEnd int,
		in []int, inStart int, inEnd int) *TreeNode
	build = func(pre []int, preStart int, preEnd int, in []int, inStart int, inEnd int) *TreeNode {
		// 前序遍历数组为空，返回 nil
		if preStart > preEnd {
			return nil
		}
		// root 节点对应的值就是前序遍历数组的第一个元素
		rootVal := pre[preStart]
		// rootVal 在中序遍历数组中的索引
		var index int
		for i := inStart; i <= inEnd; i++ {
			if in[i] == rootVal {
				index = i
				break
			}
		}
		// 构建当前节点
		root := &TreeNode{Val: rootVal}

		// 左子树的节点个数
		leftSize := index - inStart
		root.Left = build(pre, preStart+1, preStart+leftSize, in, inStart, index-1)
		root.Right = build(pre, preStart+leftSize+1, preEnd, in, index+1, inEnd)
		return root
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
