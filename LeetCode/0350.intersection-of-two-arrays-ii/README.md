# [350. 两个数组的交集Ⅱ](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)

## 题目

##### 难度：简单

给定两个数组，编写一个函数来计算它们的交集。

示例 1：

~~~
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
~~~

示例 2：

~~~
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
~~~

说明：

- 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
- 我们可以不考虑输出结果的顺序。

进阶: 

- 如果给定的数组已经排好序呢？你将如何优化你的算法？
- 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
- 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

## 解题思路

### 1.hashmap

~~~go
func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		m0[v]++
	}

	// 原地修改k
	k := 0
	for i := 0; i < len(nums2); i++ {
		if m0[nums2[i]] > 0 {
			nums2[k] = nums2[i]
			m0[nums2[i]]--
			k++
		}
	}
	return nums2[:k]
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

### 2.排序
排序后, 然后进行比较

~~~go
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j, k := 0, 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
~~~

#### 复杂度

- 时间复杂度：O(NlogN)

- 空间复杂度：O(1)
