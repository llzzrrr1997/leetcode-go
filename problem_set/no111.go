package problemset

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left == nil {
		return 1 + right
	}
	if root.Right == nil {
		return 1 + left
	}
	if left < right {
		return 1 + left
	}
	return right + 1
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MaxInt
	var f func(root *TreeNode, cnt int)
	f = func(root *TreeNode, cnt int) {
		if root.Right == nil && root.Left == nil {
			ans = min(ans, cnt)
			return
		}
		if root.Left != nil {
			f(root.Left, cnt+1)
		}
		if root.Right != nil {
			f(root.Right, cnt+1)
		}

	}
	f(root, 1)
	return ans
}
