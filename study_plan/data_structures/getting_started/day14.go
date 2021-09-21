package getting_started

//no98  中序遍历
func isValidBST(root *TreeNode) bool {
	ret := make([]int, 0)
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		traver(root.Left)
		ret = append(ret, root.Val)
		traver(root.Right)
	}
	traver(root)
	for i := 1; i < len(ret); i++ {
		if ret[i-1] >= ret[i] {
			return false
		}
	}
	return true
}

//no653
func findTarget(root *TreeNode, k int) bool {
	ret := make([]int, 0)
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		traver(root.Left)
		ret = append(ret, root.Val)
		traver(root.Right)
	}
	traver(root)
	numMap := make(map[int]bool)
	for i := 0; i < len(ret); i++ {
		if numMap[k-ret[i]] {
			return true
		}
		numMap[ret[i]] = true
	}
	return false
}

//no235
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
