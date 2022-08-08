package lc2236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNodeValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val
}

func IfEqual(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return root.Val == getNodeValue(root.Left)+getNodeValue(root.Right)
}

// -----------------------------------------------------------------------

// Another Approach:

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
