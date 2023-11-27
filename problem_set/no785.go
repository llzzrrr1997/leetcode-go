package problemset

func isBipartite(graph [][]int) bool {
	// 记录图是否符合二分图性质
	ok := true
	// 记录图中节点的颜色，false 和 true 代表两种不同颜色
	color := make([]bool, len(graph))
	// 记录图中节点是否被访问过
	visited := make([]bool, len(graph))

	// 因为图不一定是联通的，可能存在多个子图
	// 所以要把每个节点都作为起点进行一次遍历
	// 如果发现任何一个子图不是二分图，整幅图都不算二分图
	for v := 0; v < len(graph); v++ {
		if !visited[v] {
			traverse785(graph, v, visited, color, &ok)
		}
	}
	return ok
}

func traverse785(graph [][]int, v int, visited []bool, color []bool, ok *bool) {
	if !*ok {
		return
	}
	visited[v] = true
	for _, w := range graph[v] {
		if !visited[w] {
			// 相邻节点 w 没有被访问过
			// 那么应该给节点 w 涂上和节点 v 不同的颜色
			color[w] = !color[v]
			// 继续遍历 w
			traverse785(graph, w, visited, color, ok)
		} else {
			// 相邻节点 w 已经被访问过
			// 根据 v 和 w 的颜色判断是否是二分图
			if color[w] == color[v] {
				// 若相同，则此图不是二分图
				*ok = false
				return
			}
		}
	}
}
