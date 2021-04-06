package LeetCode

func removeDuplicates(nums []int) int {
	// start := 0
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			count++
			if count > 2 {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		} else {
			count = 1
		}

		// if count <= 2 {
		//     start++
		// }
	}
	// return start+1
	return len(nums)
}
