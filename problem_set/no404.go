package problemset

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lSum := sumOfLeftLeaves(root.Left)
	rSum := sumOfLeftLeaves(root.Right)

	v := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		v = root.Left.Val
	}
	return lSum + rSum + v
}
