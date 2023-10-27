package problemset

func isValidBST(root *TreeNode) bool {
	var help func(root *TreeNode, min *TreeNode, max *TreeNode) bool
	help = func(root *TreeNode, min *TreeNode, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && min.Val >= root.Val {
			return false
		}
		if max != nil && max.Val <= root.Val {
			return false
		}
		return help(root.Left, min, root) && help(root.Right, root, max)
	}
	return help(root, nil, nil)
}
