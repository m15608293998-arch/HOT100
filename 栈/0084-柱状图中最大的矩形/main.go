package main
//一句话：维护一个单调递增队列，找到当前for循环i对应的柱子的左右两边第一个比它矮的柱子
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := make([]int, 0)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			heightIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := heights[heightIndex]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			width := i - left - 1
			area := height * width
			if area > maxArea {
				maxArea = area
			}

		}
		stack = append(stack, i)
	}
	return maxArea
}

func main() {

}
