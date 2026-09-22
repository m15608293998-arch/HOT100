package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// a：头节点到环入口的距离
// b：环入口到相遇点的距离
// c：相遇点绕回入口的距离

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil

}

func main() {

}
