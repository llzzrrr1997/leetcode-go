package problemset

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var build func(nums []int, left, right int) *TreeNode
	build = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		max := -1
		index := -1
		for i := left; i <= right; i++ {
			if nums[i] > max {
				max = nums[i]
				index = i
			}
		}
		root := &TreeNode{Val: max}
		root.Left = build(nums, left, index-1)
		root.Right = build(nums, index+1, right)
		return root
	}
	return build(nums, 0, len(nums)-1)
}
