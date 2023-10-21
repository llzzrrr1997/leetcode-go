package problemset

func maxDepth104(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth104(root.Left) + 1
	right := maxDepth104(root.Right) + 1
	if left > right {
		return left
	}
	return right
}
