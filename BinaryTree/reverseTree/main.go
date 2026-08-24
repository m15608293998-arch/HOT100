package main


type TreeNode struct{
	Val int 
	Left *TreeNode
	Right *TreeNode
}

func reversetree (root *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}
	root.Left,root.Right = root.Right, root.Left
	reversetree(root.Left)
	reversetree(root.Right)
	return root
}

func main () {

}
