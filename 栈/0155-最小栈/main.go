package main

/*
 * LeetCode 155. 最小栈 (Min Stack)
 *
 * 实现思路：辅助栈法（双栈同步）
 *
 * 核心问题：普通栈只能 O(1) 拿到栈顶元素，却拿不到栈内最小值。
 * 若要在 O(1) 时间内返回最小值，就必须额外记录每个栈高度对应的最小值。
 *
 * 做法：维护两个栈
 *   - stack    主栈：正常存取所有元素；
 *   - minstack 辅助栈：与主栈保持同高度，栈顶始终保存当前栈内元素的最小值。
 *
 * 四个操作：
 *   1. Push(x)
 *      - stack 正常入栈；
 *      - minstack 压入 min(x, 当前栈顶最小值)：
 *          若 minstack 为空，直接压入 x（栈内最小值就是 x）；
 *          否则取 minstack 栈顶与 x 的较小者压入。
 *   2. Pop()
 *      - 两个栈同时弹出栈顶，保证两栈高度始终一致。
 *        主栈少一个元素后，minstack 同步弹掉该高度对应的最小值记录，
 *        栈顶自动还原为剩余元素中的最小值。
 *   3. Top()
 *      - 返回 stack 栈顶。
 *   4. GetMin()
 *      - 返回 minstack 栈顶，即当前栈内最小值。
 *
 * 正确性：minstack 的每个位置都记录了「主栈从底部到该高度为止的最小值」，
 * 所以 minstack 栈顶永远等于整个栈的最小值；弹栈时两栈同步弹出，
 * 记录随之还原，始终与栈内实际元素一一对应。
 *
 * 复杂度：
 *   - 时间复杂度：Push / Pop / Top / GetMin 均为 O(1)
 *   - 空间复杂度：O(n)，两个栈各存 n 个元素
 */

type MinStack struct {
	stack []int
	minstack []int
}
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minstack: []int{},
	}
    
}
func (this *MinStack) Push(value int)  {
	this.stack = append(this.stack, value)
	if len(this.stack) == 0{
		this.minstack = append(this.minstack, value)
		return
	}
	currentMin := this.minstack[len(this.minstack)-1]
	if value < currentMin {
		this.minstack = append(this.minstack, value)
	} else {
		this.minstack = append(this.minstack, currentMin)
	}
}

func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
    
}


func main(){
	stack := Constructor()
	stack.stack = []int{5,4}
	stack.minstack = []int{5,4}
	stack.Push(7)
	

}