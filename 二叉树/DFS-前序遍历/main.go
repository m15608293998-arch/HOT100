package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func PreOrder(root *BinaryTree) []int {
	result := make([]int, 0)
	var dfs func(node *BinaryTree)
	dfs = func(node *BinaryTree) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result

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
	fmt.Println(PreOrder(root))

}
