package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 算法思路
// 1.首先一定明确这个内部的辅助DFS在干什么：返回当前子树中是否存在p/q,如果二者都不存在那么返回nil
// 2.明晰在这个样的dfs中，边界条件是什么

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	var dfs func(node *TreeNode) *TreeNode
// 	dfs = func(node *TreeNode) *TreeNode {
// 		if node == nil {
// 			return  nil
// 		}
// 		if node == p || node == q {
// 			return node
// 		}
// 		left := dfs(node.Left)
// 		right := dfs(node.Right)
// 		if left != nil && right != nil {
// 			return node
// 		}
// 		if left != nil && right == nil {
// 			return  left
// 		}
// 		if right !=nil && left == nil {
// 			return right
// 		}
// 		return nil
// 	}
// 	return dfs(root)

// }
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
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
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right).Val)

}
