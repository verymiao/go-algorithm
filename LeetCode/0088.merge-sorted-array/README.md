# [88.Merge Sorted Array](https://leetcode-cn.com/problems/merge-sorted-array/)

## 题目

##### 难度：简单

给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。



示例 1：

~~~
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
~~~

示例 2：

~~~
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
~~~



## 解题思路

### 1.逆向双指针

合并排序数组的常用方法

~~~go
func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	for n > 0 {
		nums1[m+n-1] = nums2[n-1]
		n--
	}
}
~~~

#### 复杂度

- 时间复杂度：O(m+n)

- 空间复杂度：O(1)

### 2.其他方法

2.1 正向双指针，使用额外的空间。（略）
2.2 先合并后排序。（略）

