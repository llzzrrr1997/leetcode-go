package problemset

/*
//双重循环
func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
*/

/*//先map保存数组下表,排序后双指针(注意出现重复的数字)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int)
	for i := range nums {
		tempArr := numMap[nums[i]]
		if tempArr == nil {
			tempArr = make([]int, 0)
		}
		tempArr = append(tempArr, i)
		numMap[nums[i]] = tempArr
	}
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			leftIndex := numMap[nums[left]][0]
			numMap[nums[left]] = numMap[nums[left]][1:]
			rightIndex := numMap[nums[right]][0]
			return []int{leftIndex, rightIndex}
		}
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
	}
	return nil
}
*/

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
