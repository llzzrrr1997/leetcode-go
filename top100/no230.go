package top100

func kthSmallest(root *TreeNode, k int) int {
	var list = make([]int, 0, k)
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}
		if len(list) == k {
			return
		}
		help(root.Left)
		list = append(list, root.Val)
		help(root.Right)
	}
	help(root)

	return list[k-1]
}
