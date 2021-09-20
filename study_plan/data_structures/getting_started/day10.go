package getting_started

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//no144
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		ret = append(ret, root.Val)
		traver(root.Left)
		traver(root.Right)
	}
	traver(root)
	return ret
}

//no94
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		traver(root.Left)
		ret = append(ret, root.Val)
		traver(root.Right)
	}
	traver(root)
	return ret
}

//no145
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}
		traver(root.Left)
		traver(root.Right)
		ret = append(ret, root.Val)
	}
	traver(root)
	return ret
}
