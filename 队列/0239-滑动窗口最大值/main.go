package main
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值 。

// 算法思路：
// 1. 删除已经过期的队头
// 2. 从队尾删除所有比 nums[i] 小的元素
// 3. 当前下标 i 入队
// 4. 窗口形成后，队头就是最大值
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	ans := []int{}
	for i := 0;i < len(nums); i ++ {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:]
		}
			if len(queue) > 0 && queue[0] < i - k + 1  {
				queue = queue[1:]
			}
		queue = append(queue,i)
		if i >= k -1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return  ans
	
}

func main(){

}