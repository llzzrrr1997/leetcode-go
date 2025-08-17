package problemset

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val { //都在左子树
		return lowestCommonAncestor235(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val { //都在右子树
		return lowestCommonAncestor235(root.Right, p, q)
	}
	//在左右的子树
	return root
}
