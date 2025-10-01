package top100

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil { //在两个子树上则当前节点是祖先
		return root
	}
	if left != nil {
		return left
	}
	return right
}
