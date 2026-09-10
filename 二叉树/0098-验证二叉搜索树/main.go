package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 算法思路
// 1. 每个节点都对应一个合法取值区间 (lower, upper)，递归时把区间往下传
// 2. 左子树把上界收紧为当前值，右子树把下界收紧为当前值，一旦越界即不是 BST
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
