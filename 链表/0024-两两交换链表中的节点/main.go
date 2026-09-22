package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法思路：dummy 头 + 迭代。prev 是待交换两节点(first、second)的前驱，
// 每轮令 first.Next = second.Next、second.Next = first、prev.Next = second，然后 prev 前移到 first
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next
		first.Next = second.Next
		second.Next = first
		prev.Next = second
		prev = first
	}
	return dummy.Next

}

func main() {

}
