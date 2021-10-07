package foundation

//no199
func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		curList := []int{}
		for _, val := range tmp {
			curList = append(curList, val.Val)
			if val.Left != nil {
				q = append(q, val.Left)
			}
			if val.Right != nil {
				q = append(q, val.Right)
			}
		}
		ret = append(ret, curList[len(curList)-1])
	}
	return ret
}

//no113
func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	var helper func(root *TreeNode, leftSum int)
	helper = func(root *TreeNode, leftSum int) {
		if root == nil {
			return
		}
		leftSum -= root.Val
		path = append(path, root.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if root.Left == nil && root.Right == nil && leftSum == 0 {
			ret = append(ret, append([]int{}, path...))
		}
		helper(root.Left, leftSum)
		helper(root.Right, leftSum)
	}
	helper(root, targetSum)
	return ret
}

//no450
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	minNode := root.Right
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	root.Val = minNode.Val
	root.Right = deleteNode1(root.Right)
	return root

}

func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}
