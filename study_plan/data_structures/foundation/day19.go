package foundation

//no997
func findJudge(n int, trust [][]int) int {
	//出度为0和入度为n-1
	if n == 1 {
		return 1
	}
	cntMap := make(map[int]int)
	for _, val := range trust {
		cntMap[val[1]]++
		cntMap[val[0]]--
	}
	for key, val := range cntMap {
		if val == n-1 {
			return key
		}
	}
	return -1
}

//no1557
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	//入度为0的点
	cntMap := make([]int, n)
	for _, val := range edges {
		cntMap[val[1]]++
	}
	ret := make([]int, 0)
	for i := range cntMap {
		if cntMap[i] == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

//no841
func canVisitAllRooms(rooms [][]int) bool {
	seenNum := 0
	n := len(rooms)
	seen := make([]bool, n)
	var helper func(rooms [][]int, room int)
	helper = func(rooms [][]int, room int) {
		seen[room] = true
		seenNum++
		for _, val := range rooms[room] {
			if !seen[val] {
				helper(rooms, val)
			}
		}
	}
	helper(rooms, 0)
	return seenNum == n
}
