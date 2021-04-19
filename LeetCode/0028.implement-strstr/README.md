# [28. 实现 strStr()](https://leetcode-cn.com/problems/implement-strstr/)

## 题目

##### 难度：简单

实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。


**说明**

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

**示例 1：**

~~~
输入：haystack = "hello", needle = "ll"
输出：2
~~~

**示例 2：**

~~~
输入：haystack = "aaaaa", needle = "bba"
输出：-1
~~~

**示例 3：**

~~~
输入：haystack = "", needle = ""
输出：0
~~~

**提示：**

- 0 <= haystack.length, needle.length <= 5 * 104
- haystack 和 needle 仅由小写英文字符组成

## 解题思路

### 1.比较字符串

骚操作走一波

~~~go
func strStr(haystack string, needle string) int {
    m,n := len(haystack),len(needle)
    for i:=0;i<=m-n;i++ {
        if haystack[i:i+n] == needle {
            return i
        }
    }
    return -1
}
~~~

#### 复杂度

- 时间复杂度：O(m+n)

- 空间复杂度：O(1)

### 2.KMP算法

参考：LeetCode官方题解

~~~go
func strStr(haystack, needle string) int {
    n, m := len(haystack), len(needle)
    if m == 0 {
        return 0
    }
    pi := make([]int, m)
    for i, j := 1, 0; i < m; i++ {
        for j > 0 && needle[i] != needle[j] {
            j = pi[j-1]
        }
        if needle[i] == needle[j] {
            j++
        }
        pi[i] = j
    }
    for i, j := 0, 0; i < n; i++ {
        for j > 0 && haystack[i] != needle[j] {
            j = pi[j-1]
        }
        if haystack[i] == needle[j] {
            j++
        }
        if j == m {
            return i - m + 1
        }
    }
    return -1
}
~~~

#### 复杂度

- 时间复杂度：O(m+n)

- 空间复杂度：O(1)

