package top100

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		dif := target - v
		if j, ok := m[dif]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{}
}
