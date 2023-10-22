package problemset

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	// 1、左右子树已经被拉平成一条链表
	left := root.Left
	right := root.Right

	// 2、将左子树作为右子树
	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
