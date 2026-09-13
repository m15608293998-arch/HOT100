package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftContribution := max(0, dfs(node.Left))
		rightContribution := max(0, dfs(node.Right))
		current := leftContribution + rightContribution + node.Val
		if current > result {
			result = current
		}
		return node.Val + max(leftContribution, rightContribution)
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
	fmt.Println(maxPathSum(root))

}
