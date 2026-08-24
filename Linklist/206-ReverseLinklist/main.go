package main

import "fmt"

type Linklist struct {
	Val  int
	Next *Linklist
}

func reverselinklist(head *Linklist) *Linklist {
	if head == nil || head.Next == nil {
		return head
	}
	//栈的特性后进先出，找到尾巴节点直接return，然后进入倒数第二个节点，开始反转。
	newhead := reverselinklist(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newhead

}

func travelLinklist(head *Linklist) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}

func main() {
	head := &Linklist{
		Val: 1,
		Next: &Linklist{
			Val: 2,
			Next: &Linklist{
				Val: 3,
			},
		},
	}
	newhead := reverselinklist(head)
	travelLinklist(newhead)

}
