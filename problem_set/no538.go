package problemset

func convertBST(root *TreeNode) *TreeNode {
	//二叉树倒叙处理累加赋值
	sum := 0
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		traver(root.Right)
		sum += root.Val
		root.Val = sum
		traver(root.Left)
	}
	traver(root)
	return root
}
