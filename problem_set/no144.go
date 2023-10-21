package problemset

func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		ret = append(ret, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return ret
}
