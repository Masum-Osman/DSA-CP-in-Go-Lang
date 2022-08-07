package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Data)

	PreOrder(root.Left)
	PreOrder(root.Right)
}

func main() {
	N1 := CreateNode(1)
	N2 := CreateNode(2)
	N3 := CreateNode(3)
	N4 := CreateNode(4)
	N5 := CreateNode(5)
	N6 := CreateNode(6)
	N7 := CreateNode(7)
	N8 := CreateNode(8)

	N1.Left = N2
	N1.Right = N3

	N2.Left = N4
	N2.Right = N5

	N3.Right = N6

	N6.Left = N7
	N6.Right = N8

	fmt.Println("Pre-order traversal of Tree: ")
	PreOrder(N1)
}
