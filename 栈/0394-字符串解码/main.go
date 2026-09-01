package main

import "fmt"

func decodeString(s string) string {
	countstack := make([]int, 0)
	stringstack := make([]string, 0)
	count := 0
	current := ""
	for _, v := range s {
		if v >= '0' && v <='9' {
			count = count*10 + int(v-'0')
		} else if v == '[' {
			countstack = append(countstack, count)
			stringstack = append(stringstack, current)
			count = 0
			current = ""
		} else if v == ']' {
			repeat := countstack[len(countstack)-1]
			countstack = countstack[:len(countstack)-1]
			prev := stringstack[len(stringstack)-1]
			stringstack = stringstack[:len(stringstack)-1]
			temp := ""
			for i := 0; i < repeat; i++ {
				temp += current
			}
			current = prev + temp
		} else {
			current += string(v)
		}

	}
	return current
}

func main() {
	fmt.Println(decodeString("2a[cd]"))

}
