# [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

## 题目

##### 难度：Medium

Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.



示例 1：

~~~
输入：nums = [1,3,4,2,2]
输出：2
~~~

示例 2：

~~~
输入：nums = [3,1,3,4,2]
输出：3
~~~

示例 3：

~~~
输入：nums = [3,3,3,3,3]
输出：3
~~~

## 解题思路

### 1.暴力解法

枚举数组中的每一个数x，寻找是否存在x

~~~go
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				return nums[j]
			}
		}
	}
	return nums[0]
}

~~~

#### 复杂度

- 时间复杂度：O(N^2)

- 空间复杂度：O(1)

### 2.哈希表

使用空间换时间的思想，一次遍历

~~~go
func twoSum(nums []int, target int) []int {
    result := make([]int,0)
    hashTable := map[int]int{}
    for idx,val := range nums {
        if k,ok := hashTable[target-val]; ok {
            result = append(result,k,idx)
        }
        hashTable[val] = idx
    }
    return result
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

