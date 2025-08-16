package problemset

func sumNumbers(root *TreeNode) int {
	ret := 0

	var f func(root *TreeNode, cur int)

	f = func(root *TreeNode, cur int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			cur = cur*10 + root.Val
			ret += cur
			return
		}
		f(root.Left, cur*10+root.Val)
		f(root.Right, cur*10+root.Val)
	}
	f(root, 0)
	return ret
}
