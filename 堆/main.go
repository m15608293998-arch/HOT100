package main

// 算法思路：用数组实现的最小堆（完全二叉树）。下标 i 的左右子节点为 2i+1、2i+2，父节点为 (i-1)/2。
// Push 尾插后上浮，Pop 取走堆顶后把末尾元素搬到堆顶再下沉
type MinHeap struct {
	data []int
}

// Push：元素追加到末尾，再不断与父节点比较，比父节点小就交换（上浮）
func (h *MinHeap) Push(x int) {
	h.data = append(h.data, x)
	i := len(h.data) - 1
	for i > 0 {
		parent := (i - 1) / 2 
		if h.data[i] >= h.data[parent] {
			break
		}else {
			h.data[i] , h.data[parent] = h.data[parent] , h.data[i]
		}
		i = parent
	}

} 

// Pop：取出堆顶最小值，把末尾元素移到堆顶并缩短长度，再与较小的子节点交换（下沉）
func (h *MinHeap) Pop() int {
	 n := len(h.data)
	 minval := h.data[0]
	 h.data[0] = h.data[n-1]
	 h.data = h.data[:n-1]
	 n --
	 i := 0
	 for {
		left := 2 * i + 1
		right := 2 * i + 2
		smallest := i
		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		} 
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i] , h.data[smallest] = h.data[smallest] , h.data[i]
		i =smallest
	 }
	 return minval
}

func main(){

}