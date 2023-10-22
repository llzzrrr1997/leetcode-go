package problemset

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	var valToIdx = make(map[int]int) //存储 inorder 中值到索引的映射
	for i := 0; i < len(postorder); i++ {
		valToIdx[postorder[i]] = i
	}
	var build func(preorder []int, preStart int, preEnd int, postorder []int, postStart int, postEnd int) *TreeNode
	build = func(preorder []int, preStart int, preEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		if preStart == preEnd {
			return &TreeNode{Val: preorder[preEnd]}
		}
		// root 节点对应的值就是前序遍历数组的第一个元素
		rootVal := preorder[preStart]
		// root.left 的值是前序遍历第二个元素
		// 通过前序和后序遍历构造二叉树的关键在于通过左子树的根节点
		// 确定 preorder 和 postorder 中左右子树的元素区间
		leftRootVal := preorder[preStart+1]
		// leftRootVal 在后序遍历数组中的索引
		index := valToIdx[leftRootVal]
		// 左子树的元素个数
		leftSize := index - postStart + 1

		// 先构造出当前根节点
		root := &TreeNode{Val: rootVal}

		// 递归构造左右子树
		// 根据左子树的根节点索引和元素个数推导左右子树的索引边界
		root.Left = build(preorder, preStart+1, preStart+leftSize,
			postorder, postStart, index)
		root.Right = build(preorder, preStart+leftSize+1, preEnd,
			postorder, index+1, postEnd-1)
		return root
	}
	return build(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}
