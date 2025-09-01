package top100

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ret = append(ret, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ret
}
