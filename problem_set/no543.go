package problemset

func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	var diameter func(root *TreeNode) int
	diameter = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := diameter(root.Left)
		right := diameter(root.Right)
		cur := left + right
		if cur > ret {
			ret = cur
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	diameter(root)
	return ret
}
