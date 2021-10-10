package foundation

import (
	"container/heap"
	"strings"
)

//no451
func frequencySort(s string) string {
	h := &ArrHeap{}
	heap.Init(h)
	cntMap := make(map[int]int)
	for i := range s {
		cntMap[int(s[i])]++
	}
	for key, val := range cntMap {
		heap.Push(h, [2]int{key, val})
	}
	sb := strings.Builder{}
	for len(*h) > 0 {
		val := heap.Pop(h).([2]int)
		char := byte(val[0])
		for i := 0; i < val[1]; i++ {
			sb.WriteByte(char)
		}
	}
	return sb.String()
}

type ArrHeap [][2]int

func (h ArrHeap) Len() int           { return len(h) }
func (h ArrHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h ArrHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ArrHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *ArrHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//no973
func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}
	heap.Init(h)
	for i := 0; i < len(points); i++ {
		if i < k {
			heap.Push(h, points[i])
		} else {
			cur := heap.Pop(h).([]int)
			if (cur[0]*cur[0] + cur[1]*cur[1]) > (points[i][0]*points[i][0] + points[i][1]*points[i][1]) {
				cur = points[i]
			}
			heap.Push(h, cur)
		}
	}
	return *h
}

type PointHeap [][]int

func (h PointHeap) Len() int { return len(h) }
func (h PointHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h PointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
