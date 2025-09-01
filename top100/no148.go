package top100

import "sort"

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	valList := make([]int, 0)
	nodeList := make(map[int][]*ListNode)
	cur := head
	for cur != nil {
		if nodeList[cur.Val] != nil {
			nodeList[cur.Val] = append(nodeList[cur.Val], cur)
		} else {
			valList = append(valList, cur.Val)
			nodeList[cur.Val] = append(nodeList[cur.Val], cur)
		}
		cur = cur.Next
	}
	sort.Ints(valList)
	newHead := &ListNode{}
	cur = newHead
	for _, v := range valList {
		for _, node := range nodeList[v] {
			cur.Next = node
			cur = cur.Next
		}
	}
	cur.Next = nil

	return newHead.Next
}
