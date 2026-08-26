package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, lower, upper int64) bool
	dfs = func(node *TreeNode, lower, upper int64) bool {
		if node == nil {
			return true
		}
		val := int64(node.Val)
		if val <= lower || val >= upper {
			return false
		}
		return dfs(node.Left, lower, val) && dfs(node.Right, val, upper)

	}
	return dfs(root, -1<<63, 1<<63-1)

}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	fmt.Println(isValidBST(root))

}
