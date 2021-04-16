# [66. 加一](https://leetcode-cn.com/problems/plus-one/)

## 题目

##### 难度：简单

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。



示例 1：

~~~
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
~~~

示例 2：

~~~
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
~~~

示例 3：

~~~
输入：digits = [0]
输出：[1]
~~~


提示：

- 1 <= digits.length <= 100
- 0 <= digits[i] <= 9

## 解题思路

### 1.模拟

结尾如果不为9直接加一返回, 如果为9需要进位, 如果前面都为9数组需要增加一个1


~~~go
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(1)

