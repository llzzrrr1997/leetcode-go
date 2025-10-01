package top100

// https://leetcode.cn/problems/course-schedule/solutions/2992884/san-se-biao-ji-fa-pythonjavacgojsrust-by-pll7
func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, v := range prerequisites {
		g[v[1]] = append(g[v[1]], v[0])
	}

	colors := make([]int, numCourses)

	var dfs func(x int) bool

	dfs = func(x int) bool {
		colors[x] = 1
		for _, v := range g[x] {
			if colors[v] == 1 || (colors[v] == 0 && dfs(v)) {
				return true
			}
		}
		colors[x] = 2
		return false
	}
	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false
		}
	}

	return true
}
