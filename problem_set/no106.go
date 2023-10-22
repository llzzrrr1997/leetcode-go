package problemset

func buildTree106(inorder []int, postorder []int) *TreeNode {
	var valToIdx = make(map[int]int) //存储 inorder 中值到索引的映射
	for i := 0; i < len(inorder); i++ {
		valToIdx[inorder[i]] = i
	}
	var build func(inorder []int, inStart int, inEnd int,
		postorder []int, postStart int, postEnd int) *TreeNode
	build = func(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		// root 节点对应的值就是后序遍历数组的最后一个元素
		rootVal := postorder[postEnd]
		// rootVal 在中序遍历数组中的索引
		index := valToIdx[rootVal]
		// 左子树的节点个数
		leftSize := index - inStart
		root := &TreeNode{Val: rootVal}

		// 递归构造左右子树
		root.Left = build(inorder, inStart, index-1,
			postorder, postStart, postStart+leftSize-1)

		root.Right = build(inorder, index+1, inEnd,
			postorder, postStart+leftSize, postEnd-1)
		return root
	}
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}
