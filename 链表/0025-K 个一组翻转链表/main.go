package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// getKth 返回 prev 之后的第 k 个节点，不足 k 个则返回 nil
func getKth(prev *ListNode, k int) *ListNode {
	curr := prev
	for i := 0; i < k; i++ {
		if curr == nil {
			return nil
		}
		curr = curr.Next
	}
	return curr
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	groupPrev := dummy

	for {
		// 找当前组的第 k 个节点
		kth := getKth(groupPrev, k)
		// 不足 k 个，不翻转
		if kth == nil {
			break
		}

		groupNext := kth.Next

		// 翻转当前组
		prev := groupNext
		curr := groupPrev.Next

		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// 翻转后：
		// kth 是新头节点
		tmp := groupPrev.Next // tmp 是旧头，现在变尾节点

		groupPrev.Next = kth

		groupPrev = tmp
	}

	return dummy.Next
}

func main() {

}
