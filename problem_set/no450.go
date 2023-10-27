package problemset

func deleteNode450(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key { //找到则删除
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		//有两个子节点，为了不破坏 BST 的性质，A 必须找到左子树中最大的那个节点，或者右子树中最小的那个节点来接替自己
		maxNode := getMax(root.Left)
		root.Left = deleteNode450(root.Left, maxNode.Val)
		maxNode.Left = root.Left
		maxNode.Right = root.Right
		return maxNode
	} else if root.Val > key {
		root.Left = deleteNode450(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode450(root.Right, key)
	}
	return root
}

func getMax(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}
