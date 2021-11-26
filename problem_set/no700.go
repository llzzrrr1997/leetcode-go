package problemset

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
