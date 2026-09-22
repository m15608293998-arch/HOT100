package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 算法思路：dummy 头 + 逐位相加。双指针同步遍历两链表，sum = x + y + carry，
// 本位存 sum%10、进位记 sum/10；两链表都走完后若仍有进位再补一个节点
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		x , y := 0 , 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 =l2.Next
		}
		sum := x + y + carry
		digital := sum % 10 
		carry = sum / 10
		curr.Next = &ListNode{Val: digital}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}


func main(){

}