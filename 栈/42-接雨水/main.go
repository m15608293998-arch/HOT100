package main

// 算法思路：单调递减栈横向接水。当前柱子比栈顶高时，弹出栈顶作为坑底 bottom，
// 此时新栈顶就是左边界，宽度 = 右边界下标 - 左边界下标 - 1，高度 = min(左右边界高) - 坑底高
func trap(height []int) int {
	stack := make([]int,0)
	water := 0 
	for i:= 0; i < len(height); i ++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			width := i - left - 1
			currentwater := width * (min(height[i],height[left]) - height[bottom])
			water += currentwater
		}
		stack = append(stack, i)
	}
	return  water
	

}

func main() {

}