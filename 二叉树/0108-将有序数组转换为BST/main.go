package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST 将有序数组转换为平衡二叉搜索树（BST）。
// 思路：数组已升序，取中间元素 mid 作为根，左侧子数组递归构建左子树，
// 右侧子数组递归构建右子树。因为每次对半切分，左右子树节点数相差不超过 1，
// 所以整棵树是高度平衡的。
func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2 // 取中点，保证左右子树高度平衡
		root := &TreeNode{
			Val: nums[mid],
		}

		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)
		return root
	}
	return dfs(0, len(nums)-1)
}

func main() {
	sli := []int{1, 2, 3, 4, 5}
	fmt.Println(sortedArrayToBST(sli).Val)

}
