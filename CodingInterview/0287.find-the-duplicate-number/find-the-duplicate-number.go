package LeetCode

func findDuplicate(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}
		m[nums[i]] = struct{}{}
	}
	return nums[0]
}

func findDuplicate2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				return nums[j]
			}
		}
	}
	return nums[0]
}
