package LeetCode

func rotate(nums []int, k int) {
	n := len(nums)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[(i+k)%n] = nums[i]
	}
	copy(nums, a)
}
