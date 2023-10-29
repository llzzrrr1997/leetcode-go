package problemset

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root, root
	// 沿最左侧和最右侧分别计算高度
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	// 如果左右侧计算的高度相同，则是一棵满二叉树
	if hl == hr {
		return int(math.Pow(2, float64(hl))) - 1
	}
	// 如果左右侧的高度不同，则按照普通二叉树的逻辑计算
	return 1 + countNodes2(root.Left) + countNodes2(root.Right)
}
