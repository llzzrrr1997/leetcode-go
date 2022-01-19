package problemset

//贪心算法
func minMoves(target int, maxDoubles int) int {
	ret := 0
	for target > 1 {
		if maxDoubles > 0 {
			if target%2 == 1 { //奇数需要先减1再除2
				target--
				ret++
			}
			target = target / 2
			ret++
			maxDoubles--
			continue
		}
		target--
		ret = ret + target
		break
	}
	return ret
}
