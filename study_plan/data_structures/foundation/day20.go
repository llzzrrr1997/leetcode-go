package foundation

import "container/heap"

//no215
func findKthLargest(nums []int, k int) int {
	h := &Heap{}
	heap.Init(h)
	length := len(nums)
	for i := 0; i < length; i++ {
		if i < k {
			heap.Push(h, nums[i])
		} else {
			if nums[i] > (*h)[0] {
				heap.Pop(h)
				heap.Push(h, nums[i])
			}
		}
	}
	return (*h)[0]
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	// 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
	return h[i] < h[j]
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Pop() interface{} {
	// 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *Heap) Push(x interface{}) {
	// 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	for _, val := range nums {
		cntMap[val]++
	}
	cnt2NumsMap := make(map[int][]int)
	cntList := make([]int, 0)
	for k, v := range cntMap {
		cntList = append(cntList, v)
		cnt2NumsMap[v] = append(cnt2NumsMap[v], k)
	}
	h := &Heap{}
	heap.Init(h)
	length := len(cntList)
	for i := 0; i < length; i++ {
		if i < k {
			heap.Push(h, cntList[i])
		} else {
			if cntList[i] > (*h)[0] {
				heap.Pop(h)
				heap.Push(h, cntList[i])
			}
		}
	}
	ret := make([]int, 0)
	for _, val := range *h {
		if cnt2NumsMap[val] != nil {
			ret = append(ret, cnt2NumsMap[val]...)
			delete(cnt2NumsMap, val)
		}
	}
	return ret
}
