package LeetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
