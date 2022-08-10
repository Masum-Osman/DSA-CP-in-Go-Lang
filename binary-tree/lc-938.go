package lc938

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
 func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0

	return helper(root, low, high, &sum)
}

func helper(root *TreeNode, low int, high int, sum *int) int {
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
        *sum = *sum + root.Val
	}
	helper(root.Left, low, high, sum)
	helper(root.Right, low, high, sum)
    
    return *sum
}

// -------------------------------------------------------------

func rangeSumBST0(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	fmt.Println(low, high, root.Val)
	sum := 0
	if root.Val >= low && root.Val <= high {
		fmt.Println("Inside IF. Val: ", root.Val, "Sum: ", sum)
		sum = sum + root.Val
	}
	rangeSumBST(root.Left, low, high)
	rangeSumBST(root.Right, low, high)

	return sum
}

// ------------------------------------------------------------------

func rangeSumBST1(root *TreeNode, low int, high int) int {
	// make a slice:
	sum := 0

	helper(root , low, high, &sum)
}

func helper(root *TreeNode, low int, high int, *sum int)  {
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {

	}
	helper(root.Left, low, high, arrSum)
	helper(root.Right, low, high, arrSum)
}
