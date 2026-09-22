package main

// 算法思路：定长滑动窗口。窗口长度固定为 len(p)，用两个 26 位计数数组比较窗口内字符数是否一致；
// 右指针每扩一格，超出长度就把左指针字符挤出去
func findAnagrams(s string, p string) []int {
	ans := []int{}
	need := [26]int{}
	windows := [26]int{}
	if len(p) > len(s) {
		return ans
	}
	left := 0
	for i := 0; i < len(p); i ++ {
		need[p[i] - 'a']++
	}
	for right := 0; right < len(s); right ++ {
		windows[s[right] - 'a']++
		if right - left + 1 > len(p) {
			windows[s[left] - 'a'] --
			left ++
		}
		if right - left + 1 == len(p) && windows == need {
			ans = append(ans, left)
		}
	}
    return ans
}

func main(){

}