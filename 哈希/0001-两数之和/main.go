package main

import "fmt"

// 算法思路：哈希表一次遍历。边走边存「值→下标」，走到 v 时只需查 target-v 是否已经出现过，避免双重循环
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
