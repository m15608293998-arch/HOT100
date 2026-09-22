package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法思路：dummy 头 + 快慢指针。fast 先走 n+1 步，再与 slow 同步走到 fast 为空，
// 此时 slow 恰好停在待删节点的前驱上，执行 slow.Next = slow.Next.Next 即可
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {

}
