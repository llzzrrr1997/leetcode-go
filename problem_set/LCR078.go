package problemset

import "container/heap"

func mergeKListsLCR078(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	//最小堆实现
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	newHead := &ListNode{}
	for _, val := range lists {
		if val != nil {
			heap.Push(&pq, val)
		}
	}
	curHead := newHead
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*ListNode)
		curHead.Next = cur
		if cur.Next != nil {
			heap.Push(&pq, cur.Next)
		}
		curHead = curHead.Next
	}
	return newHead.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
