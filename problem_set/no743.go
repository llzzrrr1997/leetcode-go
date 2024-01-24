package problemset

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([][]int, 0)
	}
	for _, edge := range times {
		from, to, weight := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], []int{to, weight})
	}
	distTo := dijkstra(k, graph)
	res := 0
	for i := 1; i < len(distTo); i++ {
		if distTo[i] == math.MaxInt32 {
			//有节点不可达，返回-1
			return -1
		}
		res = max(res, distTo[i])
	}
	return res
}

func dijkstra(start int, graph [][][]int) []int {
	distTo := make([]int, len(graph))
	for i := range distTo {
		distTo[i] = math.MaxInt32
	}
	distTo[start] = 0
	q := make([][]int, 0)
	q = append(q, []int{start, 0}) //index, weight
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		if cur[1] > distTo[cur[0]] {
			continue
		}

		for _, edge := range graph[cur[0]] {
			to, weight := edge[0], edge[1]
			distToNextNode := distTo[cur[0]] + weight
			if distTo[to] > distToNextNode {
				distTo[to] = distToNextNode
				q = append(q, []int{to, distToNextNode})
			}
		}
	}
	return distTo
}
