package lc104

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)

	// if leftHeight > rightHeight {
	//     fmt.Println("left")
	//     return leftHeight + 1
	// } else {
	//     fmt.Println("right")
	//     return rightHeight + 1
	// }

	leftHeight++
	rightHeight++

	if leftHeight > rightHeight {
		fmt.Println("left")
		return leftHeight
	} else {
		fmt.Println("right")
		return rightHeight
	}

}
