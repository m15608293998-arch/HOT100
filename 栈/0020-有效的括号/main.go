package main

import (
	"fmt"
)

// func isStack(s string) bool {
// 	stack := make([]byte, 0)
// 	pairs := map[byte]byte{
// 		']': '[',
// 		')': '(',
// 		'}': '{',
// 	}
// 	for i := 0; i < len(s); i++ {
// 		ch := s[i]
// 		if ch == '[' || ch == '(' || ch == '{' {
// 			stack = append(stack, ch)
// 			continue
// 		}
// 		if len(stack) == 0 {
// 			return false
// 		}
// 		top := stack[len(stack)-1]
// 		if top != pairs[ch] {
// 			return false
// 		} else {
// 			stack = stack[:len(stack)-1]
// 		}
// 	}
// 	return len(stack) == 0

// }

// 算法思路：左括号一律入栈；遇到右括号时栈顶必须是配对的那个左括号，配对成功就弹栈；扫完后栈必须为空才算合法
func isStack(s string) bool {
	stack := make([]byte, 0)
	mapRule := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch == '{' || ch == '[' || ch == '(':
			stack = append(stack, ch)
		case ch == '}' || ch == ']' || ch == ')':
			top := stack[len(stack)-1]
			if len(stack) == 0 || top != mapRule[ch] {
				return false
			}
			stack = stack[:len(stack)-1]

		}

	}
	return len(stack) == 0
}

func main() {
	s := "[{[({})]}]"
	istrue := isStack(s)
	fmt.Println(istrue)
}
