package problemset

func searchRange(nums []int, target int) []int {
	left := leftBound(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := rightBound(nums, target)
	return []int{left, right}
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left < 0 || left >= len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target { //因为我们对 left 的更新必须是 left = mid + 1，就是说循环结束时，nums[left] 一定不等于 target 了，而 nums[left-1] 可能是 target。
			left = mid + 1
		}
	}
	if left-1 < 0 || left-1 >= len(nums) {
		return -1
	}
	if nums[left-1] == target {
		return left - 1
	}
	return -1
}
