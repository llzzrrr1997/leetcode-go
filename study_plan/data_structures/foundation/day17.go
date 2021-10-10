package foundation

//no230
func kthSmallest(root *TreeNode, k int) int {
	nums := make([]int, 0)
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		nums = append(nums, root.Val)
		helper(root.Right)
	}
	helper(root)
	return nums[k-1]
}

//no173
type BSTIterator struct {
	data []int
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	nums := make([]int, 0)
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		nums = append(nums, root.Val)
		helper(root.Right)
	}
	helper(root)
	return BSTIterator{nums}
}

func (this *BSTIterator) Next() int {
	val := this.data[0]
	this.data = this.data[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.data) > 0
}
