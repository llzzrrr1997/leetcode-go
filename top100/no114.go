package top100

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	left := root.Left
	flatten(root.Right)
	right := root.Right
	root.Left = nil
	root.Right = left
	if left == nil {
		root.Right = right
		return
	}
	if left != nil {
		for left.Right != nil { //找到最尾端的节点
			left = left.Right
		}
		left.Right = right
	}

}
