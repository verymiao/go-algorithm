# [0781.Rabbits in Forest](https://leetcode-cn.com/problems/rabbits-in-forest/)

## 题目

##### 难度：中等

森林中，每个兔子都有颜色。其中一些兔子（可能是全部）告诉你还有多少其他的兔子和自己有相同的颜色。我们将这些回答放在 answers 数组里。

返回森林中兔子的最少数量。


示例 1：

~~~
输入: answers = [1, 1, 2]
输出: 5
解释:
两只回答了 "1" 的兔子可能有相同的颜色，设为红色。
之后回答了 "2" 的兔子不会是红色，否则他们的回答会相互矛盾。
设回答了 "2" 的兔子为蓝色。
此外，森林中还应有另外 2 只蓝色兔子的回答没有包含在数组中。
因此森林中兔子的最少数量是 5: 3 只回答的和 2 只没有回答的。
~~~

示例 2：

~~~
输入: answers = [10, 10, 10]
输出: 11

输入: answers = []
输出: 0
~~~



## 解题思路

### 1.模拟

如果为0，则这个颜色兔子有一只；如果颜色B兔子有n只，至少有n+1(加上自己)；如果n+1只兔子回答n, 则第n+1兔子必然属于另一种颜色

~~~go
func numRabbits(answers []int) int {
	m := map[int]int{}
	ans := 0
	for _, v := range answers {
		if v == 0 {
			ans++
		} else if val, ok := m[v]; ok {
			val--
			if val == 0 {
				delete(m, v)
			} else {
				m[v] = val
			}
		} else {
			ans = ans + v + 1
			m[v] = v
		}
	}
	return ans
}

~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)
