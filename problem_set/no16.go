package problemset

import (
	"math"
	"sort"
)

// 16. 最接近的三数之和 https://leetcode.cn/problems/3sum-closest/solution/ji-zhi-you-hua-ji-yu-san-shu-zhi-he-de-z-qgqi/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret := 0
	minDif := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { //i-1跟i一样则跳过
			continue
		}
		//if nums[i] == nums[i+1] { //不能用这个会导致跳过完了
		//	continue
		//}

		tmp := nums[i] + nums[i+1] + nums[i+2]
		if tmp > target { //不需要再进行后续循环了，因为都比target大
			dif := int(math.Abs(float64(target - tmp)))
			if dif < minDif {
				ret = tmp
				dif = minDif
			}
			break //直接返回
		}
		tmp = nums[i] + nums[len(nums)-1] + nums[len(nums)-2]
		if tmp < target { //不需要再进行这个i的检查了，因为都比target小
			dif := int(math.Abs(float64(target - tmp)))
			if dif < minDif {
				ret = tmp
				minDif = dif
			}
			continue //跳过
		}
		left, right := i+1, len(nums)-1
		for left < right {
			tmp = nums[i] + nums[left] + nums[right]
			if tmp == target {
				return tmp
			}
			dif := int(math.Abs(float64(target - tmp)))
			if dif < minDif {
				ret = tmp
				minDif = dif
			}
			if tmp > target {
				right--
			}
			if tmp <= target {
				left++
			}
		}
	}
	return ret
}
