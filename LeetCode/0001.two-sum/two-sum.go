package LeetCode

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	hashTable := map[int]int{}
	for idx, val := range nums {
		if k, ok := hashTable[target-val]; ok {
			result = append(result, k, idx)
		}
		hashTable[val] = idx
	}
	return result
}
