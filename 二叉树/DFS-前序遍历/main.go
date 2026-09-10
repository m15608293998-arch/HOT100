package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// 算法思路：DFS 递归，按「根→左→右」的顺序访问，根在进函数时立刻记录
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
