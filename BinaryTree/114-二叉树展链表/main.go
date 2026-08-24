package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// flatten 把二叉树展平成链表，链表顺序 = 前序遍历顺序。
// 分治：先展平左右子树，再把左链移到右边、原右链接到末尾。
// 算法思路：首先看函数类型，转换成要求的树，那么少不了树的指针变动，采用分治思想，初入函数做一个基本判断
// 直接进入对于左右两子树的分支+递归，先把左边达到要求（只考虑左边先实现此效果），为了后续变量方便，将初始node的左右孩子
// 赋值为两个新的变量，开始编写，左孩子变成有右孩子，原左指针断开，此时对于整个树，为了达到需要的效果，我们定义一个tail，它代表的意义是
// 所要求右斜树的最后一个节点，但由于是分分治思想，所以当前最右边就是node节点，下一步是核心我们在上一步将左孩子变成右孩子的时候，node=tail后边
// 就已经有新的更靠右后节点了，所以开始更新，一定要找到最右后，因为我们做的还是关于左侧的所有子树的递归，会变成很深 ，最终得到新的tail
// 另一个keypoint:左边搞好了改右边，右边递归出同样要求的树，但是他没有和左边弄好的树拼接那么：tail.Right = rightchain(因为这就是右子树最顶层节点)。
func flatten(node *TreeNode) {
	if node == nil {
		return
	}

	// 1. 分治：先递归展平左右子树
	flatten(node.Left)
	flatten(node.Right)

	// 2. 取出展平后的两条链
	leftChain := node.Left
	rightChain := node.Right

	// 3. 左链挂到右边，left 置空
	node.Left = nil
	node.Right = leftChain

	// 4. 走到左链末尾
	tail := node
	for tail.Right != nil {
		tail = tail.Right
	}

	// 5. 把原来的右链接到末尾
	tail.Right = rightChain
}




func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}

	flatten(root)

	// 验证：left 全部为空，right 按前序遍历顺序串联
	for p := root; p != nil; p = p.Right {
		if p.Left != nil {
			panic("left 指针应为空")
		}
		print(p.Val, " ")
	}
	println()
}

