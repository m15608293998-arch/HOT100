package main


// 算法思路：单调递减栈存下标。当前温度高于栈顶下标对应的温度时，栈顶那天的答案就是 i - 栈顶，弹出；处理完把 i 入栈
func dailyTemperatures(temperatures []int) []int {
	result := make([]int,len(temperatures))
	stack := make([]int,0)
	for i := 0; i <len(temperatures); i ++ {
		for len(stack) > 0 && temperatures[i] > temperatures[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = i - top
		}
		stack = append(stack, i)
	}
	return result
    
}

func main(){

}