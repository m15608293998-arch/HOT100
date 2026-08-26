package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		key := string(bytes)
		groups[key] = append(groups[key], str)
	}
	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"tan", "ant", "atn", "bat", "atb", "syt"}
	fmt.Println(groupAnagrams(strs))
}
