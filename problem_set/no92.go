package problemset

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

//思路反转整个链表 再到前n个节点 再到m到n个节点。
//从前n个节点到m到n节点的反转，用相对距离看待， head索引为1，反转m, n区间， head.Next索引为1，那么就是反转m-1, n-1区间

var n1Node *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		n1Node = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = n1Node
	return last
}
