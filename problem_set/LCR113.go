package problemset

func findOrderLCR113(numCourses int, prerequisites [][]int) []int {
	graph := buildGraphLCR113(numCourses, prerequisites)
	// 计算入度，和环检测算法相同
	indegree := make([]int, numCourses)
	for _, edge := range prerequisites {
		indegree[edge[0]]++
	}
	// 根据入度初始化队列中的节点，和环检测算法相同
	q := make([]int, 0)
	for i, val := range indegree {
		if val == 0 {
			q = append(q, i)
		}
	}
	// 记录拓扑排序结果
	res := make([]int, numCourses)
	// 记录遍历节点的顺序（索引）
	count := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		// 弹出节点的顺序即为拓扑排序结果
		res[count] = cur
		count++
		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				q = append(q, next)
			}
		}
	}
	if count != numCourses {
		return []int{}
	}
	return res
}

func buildGraphLCR113(numCourses int, prerequisites [][]int) [][]int {
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
