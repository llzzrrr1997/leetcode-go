package problemset

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		// 添加一条从 from 指向 to 的有向边
		// 边的方向是「被依赖」关系，即修完课程 from 才能修课程 to
		graph[from] = append(graph[from], to)
	}
	return graph
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	// 记录一次递归堆栈中的节点
	onPath := make([]bool, numCourses)
	// 记录遍历过的节点，防止走回头路
	visited := make([]bool, numCourses)
	// 记录图中是否有环
	hasCycle := false
	for i := 0; i < numCourses; i++ {
		traverse(graph, i, &visited, &onPath, &hasCycle)
	}
	return !hasCycle
}

func traverse(graph [][]int, s int, visited *[]bool, onPath *[]bool, hasCycle *bool) {
	if (*onPath)[s] {
		// 出现环
		*hasCycle = true
		return
	}
	if (*visited)[s] || *hasCycle {
		// 如果已经找到了环，也不用再遍历了
		return
	}
	// 前序代码位置
	(*visited)[s] = true
	(*onPath)[s] = true
	for _, t := range graph[s] {
		traverse(graph, t, visited, onPath, hasCycle)
	}
	// 后序代码位置
	(*onPath)[s] = false
}
