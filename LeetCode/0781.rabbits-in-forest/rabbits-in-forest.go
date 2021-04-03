package LeetCode

// 模拟
func numRabbits(answers []int) int {
	m := map[int]int{}
	ans := 0
	for _, v := range answers {
		if v == 0 {
			ans++
		} else if val, ok := m[v]; ok {
			val--
			if val == 0 {
				delete(m, v)
			} else {
				m[v] = val
			}
		} else {
			ans = ans + v + 1
			m[v] = v
		}
	}
	return ans
}
