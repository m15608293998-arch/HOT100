package main

import "fmt"

// longestConsecutive 返回数组中最长连续序列的长度。
// 思路：先把所有数放进 set，然后只从「连续序列的起点」（num-1 不在 set 中）开始向后数。
func longestConsecutive(nums []int) int {
	mapSet := make(map[int]bool)
	for _, num := range nums {
		mapSet[num] = true
	}
	maxLength := 0

	for num := range mapSet {
		if !mapSet[num-1] {
			current := num
			length := 1

			for mapSet[current+1] {
				current++
				length++
			}
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}

func main() {
	nums := []int{100, 1, 7, 2, 3, 4}
	fmt.Println(longestConsecutive(nums))
}
