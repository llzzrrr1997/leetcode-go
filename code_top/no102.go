package code_top

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		curRet := make([]int, 0, len(curLevel))
		for i := 0; i < len(curLevel); i++ {
			curRet = append(curRet, curLevel[i].Val)
			if curLevel[i].Left != nil {
				nextLevel = append(nextLevel, curLevel[i].Left)
			}
			if curLevel[i].Right != nil {
				nextLevel = append(nextLevel, curLevel[i].Right)
			}
		}
		ret = append(ret, curRet)
		curLevel = nextLevel
	}
	return ret
}
