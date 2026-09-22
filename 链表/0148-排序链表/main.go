package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法思路：自顶向下归并排序。快慢指针找到中点并在中点前断开，递归排序左右两半，最后合并两条有序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	left := sortList(head)
	right := sortList(slow)
	return merge(left, right)
}

// merge：合并两条有序链表，返回合并后的头节点（dummy 头 + 双指针）
func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}
	if left != nil {
		curr.Next = left
		left = left.Next
	} else {
		curr.Next = right
		right = right.Next
	}
	return dummy.Next

}
func main() {

}
