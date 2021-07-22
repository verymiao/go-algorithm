# [189. 旋转数组](https://leetcode-cn.com/problems/rotate-array/)

## 题目

##### 难度：中等

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

进阶:

- 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
- 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

示例 1：

~~~
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
~~~

示例 2：

~~~
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释: 
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
~~~

提示：

- 1 <= nums.length <= 2 * 104
- -231 <= nums[i] <= 231 - 1
- 0 <= k <= 105

## 解题思路

### 1.额外数组

~~~go
func rotate(nums []int,k int) {
	n := len(nums)
	a := make([]int,n)

	for i:=0;i<n;i++ {
		a[(i+k)%n] = nums[i]
	}
	copy(nums,a)
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

### 2.反转三次

~~~go
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(1)

### 3.环状替换

~~~go
func rotate(nums []int, k int) {
    n := len(nums)
    k %= n
    for start, count := 0, gcd(k, n); start < count; start++ {
        pre, cur := nums[start], start
        for ok := true; ok; ok = cur != start {
            next := (cur + k) % n
            nums[next], pre, cur = pre, nums[next], next
        }
    }
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(1)

