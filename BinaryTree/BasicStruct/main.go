package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// 前序遍历，根左右
func PreOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// 中序遍历，左根右
func InOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)
}

// 后序遍历，左右根
func PostOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Val)
}

func main() {
	root := &BinaryTree{
		Val: 1,
		Left: &BinaryTree{
			Val: 2,
			Left: &BinaryTree{
				Val: 4,
			},
			Right: &BinaryTree{
				Val: 5,
			},
		},
		Right: &BinaryTree{
			Val: 3,
			Right: &BinaryTree{
				Val: 6,
			},
		},
	}
	fmt.Println("---前序遍历结果---")
	PreOrder(root)
	fmt.Println("---中序遍历结果---")
	InOrder(root)
	fmt.Println("---后序遍历结果---")
	PostOrder(root)

}
