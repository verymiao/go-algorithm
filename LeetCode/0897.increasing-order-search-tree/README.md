# [897. 递增顺序搜索树](https://leetcode-cn.com/problems/increasing-order-search-tree/)

## 题目

##### 难度：简单

给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。


**示例 1：**

~~~
输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
~~~

**示例 2：**

~~~
输入：root = [5,1,7]
输出：[1,null,5,null,7]
~~~

**提示：**

- 树中节点数的取值范围是 [1, 100]
- 0 <= Node.val <= 1000

## 解题思路

### 1.中序遍历之后生成新的树

~~~go
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	vals := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	dummyNode := &TreeNode{}
	curNode := dummyNode
	for _, val := range vals {
		curNode.Right = &TreeNode{Val: val}
		curNode = curNode.Right
	}
	return dummyNode.Right
}
~~~

#### 复杂度

- 时间复杂度：O(N)

- 空间复杂度：O(N)