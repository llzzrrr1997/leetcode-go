package problemset

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect(root *Node116) *Node116 {
	if root == nil {
		return nil
	}

	var traverse func(node1, node2 *Node116)
	traverse = func(node1, node2 *Node116) {
		if node1 == nil || node2 == nil {
			return
		}
		// 将传入的两个节点穿起来
		node1.Next = node2

		// 连接相同父节点的两个子节点
		traverse(node1.Left, node1.Right)
		traverse(node2.Left, node2.Right)
		// 连接跨越父节点的两个子节点
		traverse(node1.Right, node2.Left)
	}
	traverse(root.Left, root.Right)
	return root
}
