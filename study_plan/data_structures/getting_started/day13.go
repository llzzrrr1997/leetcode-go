package getting_started

//no700
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	ret := searchBST(root.Left, val)
	if ret == nil {
		return searchBST(root.Right, val)
	}
	return ret
}

//no701
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	var helper func(root *TreeNode, val int)
	helper = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if root.Val > val {
				root.Left = &TreeNode{Val: val}
				return
			}
			root.Right = &TreeNode{Val: val}
			return
		}
		if root.Left == nil && root.Val > val {
			root.Left = &TreeNode{Val: val}
			return
		}
		if root.Right == nil && root.Val < val {
			root.Right = &TreeNode{Val: val}
			return
		}
		if root.Val > val {
			helper(root.Left, val)
			return
		}
		helper(root.Right, val)
	}
	helper(root, val)
	return root
}
