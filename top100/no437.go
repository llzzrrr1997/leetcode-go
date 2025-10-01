package top100

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var ret = 0
	ret = pathSumHelper(root, targetSum)  //以root为根进行统计
	ret += pathSum(root.Left, targetSum)  //统计左子树的路径合 pathSum
	ret += pathSum(root.Right, targetSum) //统计右子树的路径合 pathSum
	return ret
}

// 从这个节点出发有多少个和为sum
func pathSumHelper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val == sum {
		ret++
	}
	ret += pathSumHelper(root.Left, sum-root.Val)
	ret += pathSumHelper(root.Right, sum-root.Val)
	return ret
}
