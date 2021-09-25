package foundation

import "sort"

//no136
func singleNumber(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}

//no169
func majorityElement(nums []int) int {
	cnt := 0
	ret := 0
	for _, val := range nums {
		if cnt == 0 {
			ret = val
		}
		if val == ret {
			cnt++
		} else {
			cnt--
		}
	}
	return ret
}

//no15
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				cur := []int{nums[i], nums[left], nums[right]}
				ans = append(ans, cur)
				for left < right && nums[left] == nums[left+1] { //排除一样的数字
					left++
				}
				for left < right && nums[right] == nums[right-1] { //排除一样的数字
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
