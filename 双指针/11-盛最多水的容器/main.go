package main

// 算法思路：左右双指针从两端向中间收缩。容量 = min(左高, 右高) × 间距，
// 每次只移动较矮的一侧——矮边才是容量的瓶颈，换掉它才可能变大
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0
	for left < right {
		h := min(height[left],height[right])
		width := right - left
		area := h * width
		if height[left] < height[right] {
			left ++
		}else {
			right --
		}
		if ans < area {
			ans = area
		}
	}
	return ans
    
}

func main(){

}