package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// 算法思路：递归求左右子树深度，取较大者 +1（空节点深度为 0）
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
// 算法思路：递归，子树和 = 根 + 左子树和 + 右子树和
func treeSum(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	leftsum := treeSum(root.Left)
	rightsum := treeSum(root.Right)
	return root.Val + leftsum + rightsum

}

// 返回树中值为 target 的节点；找不到则返回 nil
// 算法思路：递归，先查当前节点，再查左子树，最后查右子树（注意空节点要返回 nil）
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
