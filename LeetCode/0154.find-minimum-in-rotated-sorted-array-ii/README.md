# [154. 寻找旋转排序数组中的最小值 II](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)

## 题目

##### 难度：困难

已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。



示例 1：

~~~
输入：nums = [1,3,5]
输出：1
~~~

示例 2：

~~~
输入：nums = [2,2,2,0,1]
输出：0
~~~

提示：

- n == nums.length
- 1 <= n <= 5000
- -5000 <= nums[i] <= 5000
- nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转


## 解题思路

### 1.删除重复项

题目中包含排序，对于这种题目，通常采用二分查找解决。注意边界问题，可以看算题摸板里面的讲解（~~并没有写~~）

~~~go
func findMin(nums []int) int {
    lens := len(nums)
	if lens == 0 {
		return 0
	} else if lens == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
            right--
        }
	}
	return nums[left]
}
~~~

#### 复杂度

- 时间复杂度：O(lgN)

- 空间复杂度：O(1)

### 2.其他方法
- 遍历
- 排序（推荐堆排序，可以是实现接近O(N)的算法复杂度）
