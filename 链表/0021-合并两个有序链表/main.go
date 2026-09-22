package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法思路：dummy 头 + 双指针。每轮把两链表当前较小的节点接到结果尾部，
// 某一条走完后把另一条的剩余部分整体接上
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return dummy.Next

}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	head := mergeTwoLists(list1, list2)
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
	}

}
