# [1.two sum](https://leetcode-cn.com/problems/two-sum/)

## 题目

##### 难度：简单

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



示例 1：

~~~
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
~~~

示例 2：

~~~
输入：nums = [3,2,4], target = 6
输出：[1,2]
~~~



## 解题思路

### 1.暴力解法

枚举数组中的每一个数x，寻找是否存在target-x

~~~go
func twoSum(nums []int, target int) []int {
    for i, x := range nums {
        for j := i + 1; j < len(nums); j++ {
            if x+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
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

