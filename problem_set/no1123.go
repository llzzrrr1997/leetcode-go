package problemset

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	node, _ := lcaDeepestLeavesHelper(root)
	return node
}

func lcaDeepestLeavesHelper(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return root, 0
	}
	leftNode, leftDepth := lcaDeepestLeavesHelper(root.Left)
	rightNode, rightDepth := lcaDeepestLeavesHelper(root.Right)
	if leftDepth > rightDepth {
		return leftNode, leftDepth + 1
	}
	if leftDepth < rightDepth {
		return rightNode, rightDepth + 1
	}
	return root, leftDepth + 1

}
