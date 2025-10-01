package top100

func dailyTemperatures(temperatures []int) []int {
	ret := make([]int, len(temperatures))
	stack := []int{} //存下表
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			ret[j] = i - j
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ret
}
