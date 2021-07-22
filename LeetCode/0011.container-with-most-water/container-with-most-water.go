package LeetCode

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var res int
	for left < right {
		if height[left] > height[right] {
			res = max(res, height[right]*(right-left))
			right--
		} else {
			res = max(res, height[left]*(right-left))
			left++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
