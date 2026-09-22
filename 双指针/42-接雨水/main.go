package main

// 算法思路：左右双指针 + 两侧最高柱。某格的接水量由 min(左侧最高, 右侧最高) - 当前高度 决定，
// 因此每次先更新两侧最高，再结算较矮的那一侧并移动它
func trap(height []int) int {
	water := 0
	left := 0
	right := len(height) - 1
	leftmax := 0
	rightmax := 0
	for left <= right {
		if height[left] > leftmax {
			leftmax = height[left]
		}
		if height[right] > rightmax {
			rightmax = height[right]
		}
		if leftmax < rightmax {
			water += leftmax - height[left]
			left++
		} else {
			water += rightmax - height[right]
			right--
		}
	}
	return water
}

func main() {

}
