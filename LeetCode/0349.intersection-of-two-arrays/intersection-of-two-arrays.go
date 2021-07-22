package LeetCode

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	set1 := make(map[int]struct{}, 0)
	set2 := make(map[int]struct{}, 0)

	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	for v, _ := range set1 {
		if _, ok := set2[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
