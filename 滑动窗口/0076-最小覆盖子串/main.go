package main

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0

	valid := 0

	start := 0
	minLen := len(s) + 1

	for right < len(s) {
		// 1. right 字符进入窗口
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++

			if window[c] == need[c] {
				valid++
			}
		}

		// 2. 当前窗口已经覆盖 t
		//    尽可能缩小
		for valid == len(need) {

			// 更新最短答案
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			// 3. left 字符离开窗口
			d := s[left]
			left++

			if _, ok := need[d]; ok {

				if window[d] == need[d] {
					valid--
				}

				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}

func main() {

}
