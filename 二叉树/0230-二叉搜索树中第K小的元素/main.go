package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：利用二叉搜索树「中序遍历结果严格递增」的性质，第 K 小的元素就是中序遍历序列中的第 K 个节点。
// 用显式栈做中序遍历：先把根及其所有左子节点一路压栈走到最左叶子；再弹栈访问节点并将 k 减一，
// 一旦 k 归零则当前节点值即答案；随后转向该节点右子树，重复「向左压栈—弹栈计数—转右」直到找到目标或栈空。
// 时间复杂度 O(n)，空间复杂度 O(h)（h 为树高，栈保存一条根到叶路径）。
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {

			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		node = node.Right
	}
	return -1
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Println(kthSmallest(root, 3))

}
