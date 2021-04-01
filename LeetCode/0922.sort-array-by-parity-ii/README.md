# [922.按奇偶排序数组 II](https://leetcode-cn.com/problems/sort-array-by-parity-ii/)

## 题目

##### 难度：简单

给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。



示例 1：

~~~
输入：[4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
~~~



## 解题思路

### 1.两次遍历

将数组中的偶数和奇数分别取出来放进一个新的数组

~~~go
func sortArrayByParityII3(A []int) []int {
	B := make([]int, len(A))
	i := 0
	for _, v := range A {
		if v%2 == 0 {
			B[i] = v
		}
		i += 2
	}
	i = 1
	for _, v := range A {
		if v%2 == 1 {
			B[i] = v
		}
		i += 2
	}
	return B
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

### 2.一次遍历

遇到奇偶不匹配的就和后一个替换，只到匹配位置

~~~go
func sortArrayByParityII(A []int) []int {
	count := 0
	for i := 0; i < len(A); i++ {
		tmp := 1
		for A[i]%2 != count%2 {
			A[i], A[i+tmp] = A[i+tmp], A[i]
			tmp++
		}
		count++
	}
	return A
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

### 3.一次遍历, 上一个方放的改进

因为数组是奇偶各半的，不需要遍历整个数组，只需要处理所有的偶数位置是偶数，那么奇数位置就自然是奇数了

~~~go
// 只处理偶数 Best
func sortArrayByParityII2(A []int) []int {
	j := 1
	for i := 0; i < len(A); {
		for A[i]%2 != 0 {
			A[i], A[j] = A[j], A[i]
			j += 2
		}
		i += 2
	}
	return A
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)

