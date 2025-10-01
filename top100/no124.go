package top100

import "math"

func maxPathSum(root *TreeNode) int {
	ret := math.MinInt
	var help func(root *TreeNode) int

	help = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := help(root.Left)
		right := help(root.Right)
		ret = max(ret, max(left+right+root.Val, root.Val))
		return max(max(left+root.Val, right+root.Val), 0) //0比较则代表不选择这个节点
	}
	help(root)
	return ret
}
