# [剑指 Offer 03. 数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)

## 题目

##### 难度：简单

找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。



示例 1：

~~~
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3 
~~~



## 解题思路

### 1.哈希表

遇到重复就返回

~~~go
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
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

### 2.其他方法
- 暴力循环 O(N^2)
- 排序后比较 O(logN)
- 原地置换 O(N)