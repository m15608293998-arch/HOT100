package main


// 算法思路：快慢指针。fast 扫描，把非零元素依次写到 slow 位置；最后把 slow 之后的尾部全部补 0
func moveZeroes(nums []int)  {
	slow := 0
	for fast := 0; fast < len(nums); fast ++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow ++
		}
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow ++
	}
    
}

func main(){

}