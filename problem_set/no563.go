package problemset

import (
	"fmt"
	"math"
)

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	var helper func(left, right *TreeNode) (int, int)
	helper = func(left, right *TreeNode) (int, int) {
		leftSum := 0
		rightSum := 0
		leftRet := 0
		rightRet := 0
		leftVal := 0
		rightVal := 0
		sum := 0
		if left != nil {
			leftSum, leftRet = helper(left.Left, left.Right)
			leftVal = left.Val
		}
		if right != nil {
			rightSum, rightRet = helper(right.Left, right.Right)
			rightVal = right.Val
		}
		ret += leftRet + rightRet
		sum += leftSum + rightSum + leftVal + rightVal
		dif := int(math.Abs(float64(leftSum+leftVal) - float64(rightSum+rightVal)))
		fmt.Println(sum, dif)
		return sum, dif
	}
	_, rootRet := helper(root.Left, root.Right)
	ret += rootRet
	return ret
}
