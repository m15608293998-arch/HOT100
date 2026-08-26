package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func maxdepth(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	leftdepth := maxdepth(root.Left)
	rightdepth := maxdepth(root.Right)
	if leftdepth > rightdepth {
		return leftdepth + 1
	}

	return rightdepth + 1

}

// 返回树中所有节点的值
func treeSum(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	leftsum := treeSum(root.Left)
	rightsum := treeSum(root.Right)
	return root.Val + leftsum + rightsum

}

// 返回树中值为 target 的节点；找不到则返回 nil
func findNode(root *BinaryTree, target int) *BinaryTree {
	if root.Val == target {
		return root
	} else {
		return nil
	}
	lefnode := findNode(root.Left, target)
	if lefnode != nil {
		return lefnode
	}
	return findNode(root.Right, target)

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
	fmt.Println(maxdepth(root))
	fmt.Println(treeSum(root))

}
