package problemset

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left == nil {
		return 1 + right
	}
	if root.Right == nil {
		return 1 + left
	}
	if left < right {
		return 1 + left
	}
	return right + 1
}
