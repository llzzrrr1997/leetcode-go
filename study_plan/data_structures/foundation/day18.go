package foundation

import (
	"fmt"
	"strconv"
	"strings"
)

//no236
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := strings.Builder{}
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val) + ",")
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	ret := sb.String()
	ret = ret[:len(ret)-1]
	fmt.Println(ret)
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	fmt.Println(len(sp))
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}
