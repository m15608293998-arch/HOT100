package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return nil
}

func main() {
	nums := []int{1, 3, 4, 5, 9}
	target := 9
	fmt.Println(twoSum(nums, target))
}
