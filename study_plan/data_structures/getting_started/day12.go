package getting_started

//no226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

//no112
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	dif := targetSum - root.Val
	if dif == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	if dif == 0 && root.Left == nil && root.Right == nil {
		return false
	}
	return hasPathSum(root.Left, dif) || hasPathSum(root.Right, dif)
}
