# [15. 三数之和](https://leetcode-cn.com/problems/3sum/)

## 题目

##### 难度：中等

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

~~~
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
~~~

示例 2：

~~~
输入：nums = []
输出：[]
~~~

示例 3：

~~~
输入：nums = [0]
输出：[]
~~~

提示：

- 0 <= nums.length <= 3000
- -105 <= nums[i] <= 105


## 解题思路

### 1.

~~~go
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		right := len(nums) - 1
		for left := i + 1; left < len(nums); left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}
			for left < right && nums[left]+nums[right] > target {
				right--
			}

			if left == right {
				break
			}
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
			}
		}
	}
	return result
}
~~~

#### 复杂度

- 时间复杂度：O(N^2)

- 空间复杂度：O(lgN)

### 2.

~~~go
func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] < target {
				left++
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return result
}
~~~

- 时间复杂度：O(N^2)

- 空间复杂度：O(lgN)