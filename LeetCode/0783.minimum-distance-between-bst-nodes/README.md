# [783. 二叉搜索树节点最小距离](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/)

## 题目

##### 难度：简单

给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同



示例 1：

~~~
输入：root = [4,2,6,1,3]
输出：1
~~~


示例 2：

~~~
输入：root = [1,0,48,null,null,12,49]
输出：1
~~~

## 解题思路

### 1.递归

中序遍历，递归。常用解法

~~~go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    pre, diff :=-1, 1 << 62

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return 
        }
        dfs(node.Left)
        if pre != -1 && node.Val - pre < diff {
            diff = node.Val - pre
        }
        pre = node.Val
        dfs(node.Right)
    }
    dfs(root)
    return diff
}

~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(1)

### 2.迭代遍历

参考链接: https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/solution/go-shuang-100-zhong-xu-bian-li-die-dai-fang-shi-by/

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(1)