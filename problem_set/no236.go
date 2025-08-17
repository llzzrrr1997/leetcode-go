package problemset

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q { //当前节点是空、q、p都是返回当前节点
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)   //左子树的结果
	right := lowestCommonAncestor(root.Right, p, q) //右子树的结果
	if left != nil && right != nil {                //如果左右子树都不为空，证明p、q在左右子树，当前root就是结果
		return root
	}
	if left != nil { //p、q都在左子树，返回左子树
		return left
	}
	return right //都在右子树，返回右子树
}
