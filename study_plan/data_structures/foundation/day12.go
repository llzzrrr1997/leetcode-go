package foundation

//no24
func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	cur := newHead
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		cur.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		cur = node1
	}
	return newHead.Next
}

//no707

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	size int
	head *Node
}

func MyLinkedListConstructor() MyLinkedList {
	return MyLinkedList{0, &Node{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index <= 0 {
		index = 0
	}
	this.size++
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	node := &Node{val, cur.Next}
	cur.Next = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
