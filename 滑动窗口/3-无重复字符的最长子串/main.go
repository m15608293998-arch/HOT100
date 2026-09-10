package main

// 算法思路：滑动窗口 + 哈希表记录每个字符最后出现的下标。
// 右指针遇到重复字符时，左指针直接跳到该字符上次出现的下一位（跳跃式收缩，不用一格一格挪）
func lengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	left := 0 
	maxlen := 0 
	for right :=0 ; right < len(s); right ++ {
		ch := s[right] 
		if pos,ok := last[ch]; ok && pos >= left {
			left = pos + 1 
		}
		last[ch] = right
		len := right - left + 1 
		if len > maxlen	{
			maxlen = len
		}
	}
	return  maxlen
}

// func lengthOfLongestSubstring(s string) int {
// 	count := make(map[byte]int)
// 	left := 0
// 	maxlen := 0
// 	for right := 0; right < len(s); right ++ {
// 		count[s[right]]++
// 		for count[s[right]] > 1 {
// 			count[s[left]]--
// 			left ++
// 		}
// 		if right - left + 1 >  maxlen {
// 			maxlen = right - left + 1
// 		}
// 	}
// 	return maxlen
    
// }

func main(){

}