package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int, len(inorder))
	for k, v := range inorder {
		indexMap[v] = k
	}
	preindex := 0
	var dfs func(inleft, inright int) *TreeNode
	dfs = func(inleft, inright int) *TreeNode {
		if inleft > inright {
			return nil
		}
		rootVal := preorder[preindex]
		preindex++
		root := &TreeNode{
			Val: rootVal,
		}
		rootIndex := indexMap[rootVal]
		root.Left = dfs(inleft, rootIndex-1)
		root.Right = dfs(rootIndex+1, inright)
		return root
	}
	return dfs(0, len(inorder)-1)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root)

}
