# [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

## 题目

##### 难度：简单

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例：

~~~
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
~~~

说明：

- 1.必须在原数组上操作，不能拷贝额外的数组。
- 2.尽量减少操作次数。

## 解题思路

### 1.双指针
快慢指针

~~~go
func moveZeroes(nums []int)  {
	j:=0 // 慢指针
	for i:=0;i<len(nums) ;i++  {
		if nums[i] !=0 {
			if i!=j {
				nums[i],nums[j] = nums[j],nums[i]
			}
			j++
		}
	}
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(1)
