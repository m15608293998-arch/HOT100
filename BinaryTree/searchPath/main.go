package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func hasPathSum(root *BinaryTree, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	remain := targetSum - root.Val
	return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
func main() {
	root := &BinaryTree{
		Val: 5,
		Left: &BinaryTree{
			Val: 10,
			Left: &BinaryTree{
				Val: 4,
			},
			Right: &BinaryTree{
				Val: 6,
			},
		},
		Right: &BinaryTree{
			Val: 3,
			Right: &BinaryTree{
				Val: 6,
			},
		},
	}
	fmt.Println(hasPathSum(root, 21))

}
