package problemset

func possibleBipartition(n int, dislikes [][]int) bool {
	ok := true
	color := make([]bool, n+1)
	visited := make([]bool, n+1)
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, val := range dislikes {
		graph[val[0]] = append(graph[val[0]], val[1])
		graph[val[1]] = append(graph[val[1]], val[0])
	}
	for v := 1; v <= n; v++ {
		if !visited[v] {
			traverse886(graph, v, &visited, &color, &ok)
		}
	}
	return ok
}

func traverse886(graph [][]int, v int, visited *[]bool, color *[]bool, ok *bool) {
	if !*ok {
		return
	}
	(*visited)[v] = true
	for _, w := range graph[v] {
		if !(*visited)[w] {
			(*color)[w] = !(*color)[v]
			traverse886(graph, w, visited, color, ok)
		} else {
			if (*color)[w] == (*color)[v] {
				*ok = false
				return
			}
		}
	}
}
