package LeetCode

func findRepeatNumber(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = 1
	}
	return -1
}
