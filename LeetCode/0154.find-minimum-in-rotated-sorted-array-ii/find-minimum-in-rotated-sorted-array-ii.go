package LeetCode

func findMin(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	} else if lens == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
