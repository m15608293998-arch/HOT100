package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 算法思路：递归，每个节点交换自己的左右孩子，再递归处理两棵子树
func reversetree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	reversetree(root.Left)
	reversetree(root.Right)
	return root
}

func main() {

}
