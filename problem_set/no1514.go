package problemset

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	graph := make([][]tuple, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]tuple, 0)
	}
	for i, edge := range edges {
		from := edge[0]
		to := edge[1]
		weight := succProb[i]
		graph[from] = append(graph[from], tuple{to, weight})
		graph[to] = append(graph[to], tuple{from, weight})
	}
	probTo := make([]float64, len(graph))
	probTo[start_node] = 1.0
	q := make([]state, 0)
	q = append(q, state{id: start_node, probFromStart: 1.0})
	for len(q) > 0 {
		curState := q[0]
		q = q[1:]
		curNodeID := curState.id
		curProbFromStart := curState.probFromStart
		if probTo[curNodeID] > curProbFromStart {
			continue
		}
		for _, neighbor := range graph[curNodeID] {
			nextNodeID := neighbor.to
			// 看看从 curNode 达到 nextNode 的概率是否会更大
			probToNextNode := probTo[curNodeID] * neighbor.prob
			if probTo[nextNodeID] < probToNextNode {
				probTo[nextNodeID] = probToNextNode
				q = append(q, state{
					nextNodeID,
					probToNextNode,
				})
			}
		}
	}
	return probTo[end_node]
}

type state struct {
	// 图节点的 id
	id int
	// 从 start 节点到达当前节点的概率
	probFromStart float64
}

type tuple struct {
	to   int
	prob float64
}
