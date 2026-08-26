package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 算法思路
// 1.通过DFS从上往下走传递状态
// 2.判断对称状态
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(left, right *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return dfs(left.Right, right.Left) && dfs(left.Left, right.Right)
	}
	return dfs(root.Left, root.Right)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(root))
}
