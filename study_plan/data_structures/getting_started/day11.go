package getting_started

//no102
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ret := make([][]int, 0)
	for len(queue) > 0 {
		cnt := len(queue)
		curRet := make([]int, 0, cnt)
		tempQueue := make([]*TreeNode, 0)
		for i := 0; i < cnt; i++ {
			curRet = append(curRet, queue[i].Val)
			if queue[i].Left != nil {
				tempQueue = append(tempQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tempQueue = append(tempQueue, queue[i].Right)
			}
		}
		queue = tempQueue
		ret = append(ret, curRet)
	}
	return ret
}

//104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var helper func(root *TreeNode, cur int) int
	helper = func(root *TreeNode, cur int) int {
		if root == nil {
			return cur
		}
		leftDepth := helper(root.Left, cur) + 1
		rightDepth := helper(root.Right, cur) + 1
		return max(leftDepth, rightDepth)
	}
	return helper(root, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//np101
func isSymmetric(root *TreeNode) bool {
	var helper func(left, right *TreeNode) bool
	helper = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left != nil && right == nil {
			return false
		}
		if left == nil && right != nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return helper(left.Left, right.Right) && helper(left.Right, right.Left)
	}
	return helper(root.Left, root.Right)
}
