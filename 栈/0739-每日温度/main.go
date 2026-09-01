package main


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