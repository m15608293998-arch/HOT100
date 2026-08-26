package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 算法思路
// 1.直径 = 当前节点最大左子树高度 + 最大右子树高度，求高度用 DFS
// 2.核心：最大直径不一定经过 root，所以需要在每一轮 DFS 中用一个 result 记录最大直径
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		diameter := left + right
		if diameter > result {
			result = diameter
		}
		return max(left, right) + 1
	}
	dfs(root)
	return result

}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(diameterOfBinaryTree(root))

}
