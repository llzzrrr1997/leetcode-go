package problemset

//二叉树遍历 递归
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		ret = append(ret, root.Val)
		helper(root.Right)
	}
	helper(root)
	return ret
}

//二叉树遍历 迭代
func inorderTraversal2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	for len(stack) > 0 || root != nil {
		for root != nil { //左子树遍历完
			stack = append(stack, root)
			root = root.Left
		}
		n := len(stack)
		ret = append(ret, stack[n-1].Val)
		root = stack[n-1].Right
		stack = stack[:n-1]
	}
	return ret
}
