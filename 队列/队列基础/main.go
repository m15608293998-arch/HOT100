package main

import "fmt"

func main() {
	queue := make([]int, 0)
	queue = append(queue, 10)
	queue = append(queue, 20)
	queue = append(queue, 30)
	if len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf("出队:%d\n", front)
	}
	fmt.Println(queue)

}
