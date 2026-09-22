package main

import "sort"

// 算法思路：先排序，枚举第一个数 nums[i]，剩余两数用左右双指针夹逼（和偏小左移、偏大右移）。
// 排序后相同元素相邻，靠跳过重复值完成去重
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums) - 2 ; i ++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1 
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left ++
			}else if sum > 0{
				right --
			}else {
				ans = append(ans, []int{nums[i],
					nums[left],nums[right],
				})
				left ++
				right --
				for left < right && nums[left] == nums[left-1] {
					left ++
				}
				for left < right && nums[right] == nums[right+1] {
					right --
				}
			}
		}

	}
	return ans

}

func main() {

}