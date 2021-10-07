package foundation

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//no108
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(data []int, left, right int) *TreeNode
	helper = func(data []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{Val: data[mid]}
		root.Left = helper(data, left, mid-1)
		root.Right = helper(data, mid+1, right)
		return root
	}
	return helper(nums, 0, len(nums)-1)
}

//no105
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := 0
	for ; index < len(inorder); index++ {
		if preorder[0] == inorder[index] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}

//no103
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	level := 1
	for len(queue) > 0 {
		tmp := make([]int, 0)
		q := queue
		queue = nil
		for _, node := range q {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 0 {
			for i, n := 0, len(tmp); i < n/2; i++ {
				tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
			}
		}
		level++
		ret = append(ret, tmp)
	}
	return ret
}
