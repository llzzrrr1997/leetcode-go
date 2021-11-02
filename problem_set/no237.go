package problemset

func deleteNode(node *ListNode) {
	//把下个节点的值赋值给当前节点
	//把当前节点的next指向下下个节点，相当于删除下个节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
