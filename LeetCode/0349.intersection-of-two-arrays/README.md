# [349. 两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays/)

## 题目

##### 难度：简单

给定两个数组，编写一个函数来计算它们的交集。

示例 1：

~~~
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
~~~

示例 2：

~~~
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
~~~

说明：

- 输出结果中的每个元素一定是唯一的。
- 我们可以不考虑输出结果的顺序。

## 解题思路

### 1.hashmap
使用两个map存储两个数组, 然后比较map

~~~go
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
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

### 2.排序
排序后, 然后进行比较

~~~go
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if res == nil || res[len(res)-1] != nums1[i] {
				res = append(res, nums2[j])
			}
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else if nums2[j] > nums1[i] {
			i++
		}
	}
	return res
}
~~~

#### 复杂度

- 时间复杂度：O(NlogN)

- 空间复杂度：O(1)
