package problemset

func kthSmallest(root *TreeNode, k int) int {
	var ret int
	var traver func(root *TreeNode, k int, rank *int)
	traver = func(root *TreeNode, k int, rank *int) {
		if root == nil {
			return
		}
		traver(root.Left, k, rank)
		*rank++
		if *rank == k {
			ret = root.Val
			return
		}
		traver(root.Right, k, rank)
	}
	rank := 0
	traver(root, k, &rank)
	return ret
}
