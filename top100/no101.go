package top100

func isSymmetric(root *TreeNode) bool {
	var help func(right, left *TreeNode) bool

	help = func(right, left *TreeNode) bool {

		if right == nil && left == nil {
			return true
		}

		if right != nil && left == nil {
			return false
		}
		if right == nil && left != nil {
			return false
		}
		if right.Val != left.Val {
			return false
		}

		return help(right.Left, left.Right) && help(right.Right, left.Left)

	}
	if root == nil {
		return true
	}
	return help(root.Left, root.Right)
}
