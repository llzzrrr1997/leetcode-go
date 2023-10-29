package problemset

func allPathsSourceTarget(graph [][]int) [][]int {
	ret := [][]int{}
	curPath := []int{}
	var help func(graph [][]int, curPath *[]int, ret *[][]int, s int)
	help = func(graph [][]int, curPath *[]int, ret *[][]int, s int) {
		*curPath = append(*curPath, s)
		n := len(graph)
		if s == n-1 {
			*ret = append(*ret, append([]int{}, *curPath...))
			*curPath = (*curPath)[:len(*curPath)-1]
			return
		}
		for _, val := range graph[s] {
			help(graph, curPath, ret, val)
		}
		*curPath = (*curPath)[:len(*curPath)-1]
	}
	help(graph, &curPath, &ret, 0)
	return ret
}

func allPathsSourceTarget2(graph [][]int) [][]int {
	ret := [][]int{}
	curPath := []int{}
	var help func(s int)
	help = func(s int) {
		curPath = append(curPath, s)
		n := len(graph)
		if s == n-1 {
			ret = append(ret, append([]int{}, curPath...))
			curPath = (curPath)[:len(curPath)-1]
			return
		}
		for _, val := range graph[s] {
			help(val)
		}
		curPath = (curPath)[:len(curPath)-1]
	}
	help(0)
	return ret
}

//// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
//// 本代码不保证正确性，仅供参考。如有疑惑，可以参照我写的 java 代码对比查看。
//
//// 记录被遍历过的节点
//var visited []bool
//// 记录从起点到当前节点的路径
//var onPath []bool
//
///* 图遍历框架 */
//func traverse(graph Graph, s int) {
//    if visited[s] {
//        return
//    }
//    // 经过节点 s，标记为已遍历
//    visited[s] = true
//    // 做选择：标记节点 s 在路径上
//    onPath[s] = true
//    for _, neighbor := range graph.neighbors(s) {
//        traverse(graph, neighbor)
//    }
//    // 撤销选择：节点 s 离开路径
//    onPath[s] = false
//}
