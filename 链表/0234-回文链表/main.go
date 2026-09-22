package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法思路：快慢指针找到中点，反转后半段，再与前半段逐个比对。
// 时间 O(n)、额外空间 O(1)，代价是会就地修改后半段的指针方向
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	L1 := head
	L2 := prev
	for L2 != nil {
		if L1.Val != L2.Val {
			return false
		}
		L1 = L1.Next
		L2 = L2.Next
	}
	return true
}

func main() {

}
