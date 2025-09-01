package top100

func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	var help func(root *TreeNode) int

	help = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := help(root.Left)
		right := help(root.Right)
		if root.Left != nil {
			left += 1
		}
		if root.Right != nil {
			right += 1
		}
		ret = max(ret, left+right)
		return max(left, right)
	}
	help(root)
	return ret
}
