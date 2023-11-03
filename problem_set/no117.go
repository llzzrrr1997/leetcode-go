package problemset

func connect117(root *Node116) *Node116 {
	//非完美二叉树，使用层序遍历然后进行赋值
	if root == nil {
		return nil
	}
	q := []*Node116{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i := 0; i < len(tmp); i++ {
			if i+1 <= len(tmp) {
				tmp[i].Next = tmp[i+1]
			}
			if tmp[i].Left != nil {
				q = append(q, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				q = append(q, tmp[i].Right)
			}
		}
	}
	return root
}

func connect117_2(root *Node116) *Node116 {
	//非完美二叉树，使用上一层的串联就能连接下一层的节点
	start := root
	for start != nil {
		var nextStart, lastNode *Node116
		help := func(cur *Node116) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if lastNode != nil {
				lastNode.Next = cur
			}
			lastNode = cur
		}
		for p := start; p != nil; p = p.Next {
			help(p.Left)
			help(p.Right)
		}
		start = nextStart
	}
	return root
}
