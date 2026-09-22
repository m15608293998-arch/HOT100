package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法思路：快慢指针（Floyd 判圈）。slow 每次走 1 步、fast 每次走 2 步，
// 有环则 fast 必在环内追上 slow；fast 先走到尽头说明无环
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false

}

func main() {

}
