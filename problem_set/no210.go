package problemset

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 后序遍历结果（拓扑排序就是后序遍历的逆序）
	postorder := []int{}
	// 是否存在环
	hasCycle := false
	// 节点访问状态
	visited := make([]bool, numCourses)
	// 节点遍历状态
	onPath := make([]bool, numCourses)

	// 构建图
	graph := buildGraph(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		traverse210(graph, i, &postorder, &hasCycle, visited, onPath)
	}
	if hasCycle {
		return []int{}
	}
	reverse210(postorder)
	return postorder
}

func traverse210(graph [][]int, s int, postorder *[]int, hasCycle *bool, visited []bool, onPath []bool) {
	if onPath[s] {
		*hasCycle = true
	}
	if visited[s] || *hasCycle {
		return
	}
	onPath[s] = true
	visited[s] = true
	for _, t := range graph[s] {
		traverse210(graph, t, postorder, hasCycle, visited, onPath)
	}
	onPath[s] = false
	*postorder = append(*postorder, s)
}
func reverse210(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
