package main

import "fmt"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// BFS层序遍历
func levleOrder(root *BinaryTree) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := []*BinaryTree{root}

	for len(queue) > 0 {
		levelsieze := len(queue)
		level := make([]int, 0, levelsieze)
		for i := 0; i < levelsieze; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result

}

// BFS引申，看一颗二叉树的右视图
func rightWatch(root *BinaryTree) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	queue := []*BinaryTree{root}
	for len(queue) > 0 {
		levelsize := len(queue)
		level := make([]int, 0, levelsize)
		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == levelsize-1 {
				level = append(level, node.Val)
			}
		}
		result = append(result, level...)
	}
	return result
}

func levelSums(root *BinaryTree) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	queue := []*BinaryTree{root}
	for len(queue) > 0 {
		levelsize := len(queue)
		sum := 0
		for i := 0; i < levelsize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, sum)
	}
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
	fmt.Println(levleOrder(root))
	fmt.Println(rightWatch(root))
	fmt.Println(levelSums(root))

}
