package LeetCode

func sortArrayByParityII(A []int) []int {
	j := 1
	for i := 0; i < len(A); {
		for A[i]%2 != 0 {
			for A[j]%2 == 0 {
				A[i], A[j] = A[j], A[i]
			}
			j += 2
		}
		i += 2
	}
	return A
}
