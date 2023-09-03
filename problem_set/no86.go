package problemset

import "fmt"

func partition(head *ListNode, x int) *ListNode {
	var smallHead, smallTail, bigHead, bigTail *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		fmt.Println(tmp.Val)
		if tmp.Val < x {
			if smallHead == nil {
				smallHead = tmp
				smallTail = tmp
			} else {
				smallTail.Next = tmp
				smallTail = smallTail.Next
			}
		} else {
			if bigHead == nil {
				bigHead = tmp
				bigTail = tmp
			} else {
				bigTail.Next = tmp
				bigTail = bigTail.Next
			}
		}

	}
	fmt.Println(smallHead)
	if smallHead == nil {
		return bigHead
	}
	if bigHead == nil {
		return smallHead
	}
	bigTail.Next = nil //设置为空
	smallTail.Next = bigHead
	return smallHead
}
