package main

import "fmt"

type BST struct {
	Val   int
	Left  *BST
	Right *BST
}

var count int

func (root *BST) InsertBST(data int) {
	if root.Val < data {
		//move right
		if root.Right == nil {
			root.Right = &BST{Val: data}
		} else {
			root.Right.InsertBST(data)
		}
	} else if root.Val > data {
		//move left
		if root.Left == nil {
			root.Left = &BST{Val: data}
		} else {
			root.Left.InsertBST(data)
		}
	}
}

func (root *BST) IfContainsBST(data int) bool {
	count++

	if root == nil {
		return false
	}

	if data > root.Val {
		//look left
		return root.Left.IfContainsBST(data)
	} else if data < root.Val {
		//look right
		return root.Right.IfContainsBST(data)
	}

	return true
}

func BSTTraversal(root *BST) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	BSTTraversal(root.Left)
	BSTTraversal(root.Right)
}

func main() {
	bst := &BST{Val: 100}

	bst.InsertBST(50)
	bst.InsertBST(45)
	bst.InsertBST(152)
	bst.InsertBST(14)
	bst.InsertBST(545)
	bst.InsertBST(4)
	bst.InsertBST(54)
	bst.InsertBST(45)
	bst.InsertBST(7)
	bst.InsertBST(47)
	bst.InsertBST(9)
	bst.InsertBST(87)
	bst.InsertBST(36)
	bst.InsertBST(263)
	bst.InsertBST(89)
	bst.InsertBST(636)
	bst.InsertBST(66)
	bst.InsertBST(78)
	bst.InsertBST(6)

	BSTTraversal(bst)
	fmt.Println("\nTraversing Done")
	fmt.Println(bst.IfContainsBST(66))
	fmt.Println(count)
}
