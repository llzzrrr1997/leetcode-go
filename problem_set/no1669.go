package problemset

import "fmt"

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	headNode2 := list2
	tailNode2 := list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	tailNode2 = list2
	preANode := list1
	for a > 1 {
		a--
		preANode = preANode.Next
	}
	bNode := list1
	for b > 0 {
		bNode = bNode.Next
		b--
	}
	fmt.Println(headNode2.Val)
	fmt.Println(tailNode2.Val)
	fmt.Println(preANode.Val)
	fmt.Println(bNode.Val)
	preANode.Next = headNode2
	tailNode2.Next = bNode.Next
	return list1
}
