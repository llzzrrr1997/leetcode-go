package problemset

import "math"

func isValidBST(root *TreeNode) bool {
	var help func(root *TreeNode, min *TreeNode, max *TreeNode) bool
	help = func(root *TreeNode, min *TreeNode, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && min.Val >= root.Val {
			return false
		}
		if max != nil && max.Val <= root.Val {
			return false
		}
		return help(root.Left, min, root) && help(root.Right, root, max)
	}
	return help(root, nil, nil)
}

func isValidBST2(root *TreeNode) bool {
	var check func(rooT *TreeNode, minVal, maxVal int) bool

	check = func(root *TreeNode, minVal, maxVal int) bool {
		if root == nil {
			return true
		}
		return root.Val > minVal && root.Val < maxVal && check(root.Left, minVal, root.Val) && check(root.Right, root.Val, maxVal)
	}
	return check(root, math.MinInt, math.MaxInt)
}
