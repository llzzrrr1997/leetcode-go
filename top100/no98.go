package top100

import "math"

func isValidBST(root *TreeNode) bool {
	var helper func(root *TreeNode, minVal, maxVal int) bool
	helper = func(root *TreeNode, minVal, maxVal int) bool {
		if root == nil {
			return true
		}
		if !(minVal < root.Val && root.Val < maxVal) {
			return false
		}
		return helper(root.Left, minVal, root.Val) && helper(root.Right, root.Val, maxVal)
	}
	return helper(root, math.MinInt, math.MaxInt)
}
