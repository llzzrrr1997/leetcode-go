package getting_started

//no1
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dif := target - nums[i]
		j, ok := numMap[dif]
		if ok {
			return []int{i, j}
		}
		numMap[nums[i]] = i
	}
	return nil
}

/*func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	for i := 0; i < m; i++ {
		temp[i] = nums1[i]
	}
	i, j, k := 0, 0, 0
	for ; i < m || j < n; k++ {
		if i < m && j < n {
			if temp[i] > nums2[j] {
				nums1[k] = nums2[j]
				j++
			} else {
				nums1[k] = temp[i]
				i++
			}
			continue
		}
		if i < m {
			nums1[k] = temp[i]
			i++
		}
		if j < n {
			nums1[k] = nums2[j]
			j++
		}
	}
}*/

//no88
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	for k := m + n - 1; k >= 0; k-- {
		if j < 0 {
			continue
		}
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
