package top100

func majorityElement(nums []int) int {
	ret := 0
	cnt := 0
	for _, val := range nums {
		if cnt == 0 {
			cnt++
			ret = val
		} else if ret == val {
			cnt++
		} else if ret != val {
			cnt--
		}
	}
	return ret
}
