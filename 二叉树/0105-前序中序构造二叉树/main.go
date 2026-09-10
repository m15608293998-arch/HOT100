package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 算法思路
// 1. 前序的首个元素是根；用哈希表把中序的「值→下标」映射，O(1) 定位根的位置
// 2. 中序中根左边是左子树、右边是右子树，据此划分区间递归构造
// 3. preindex 随递归顺序全局递增，保证每层取到的根都符合前序顺序
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
