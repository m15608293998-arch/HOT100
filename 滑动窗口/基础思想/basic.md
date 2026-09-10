
滑动窗口其实没有单调栈那么“技巧性强”，它的核心就一句话：每个窗口维护题目所需要的数据，用hash map，快速取用
用两个指针维护一个连续区间，并利用“窗口移动后只变化少量元素”这一点，避免每次重新计算整个区间，一个连续窗口就需要双指针[*left,*right]
两类题目：1.固定长度窗口：长度固定为 k，求每个窗口的最大值、和、平均值 2.可变长度滑动窗口：根据窗口是否合法，动态扩大或缩小。
滑动窗口的滑动逻辑：什么是不变量：比如串中最长不重复子串，就是必须让窗口满足不重复，如果满足right++，不满足那么left++
1. 模板
func pattern1(nums []int, k int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		//nums[right]进入窗口
		if right-left+1 > k {
			//nums[left]离开窗口
			left++
		}
		if right-left+1 == k {
			//更新答案
		}
	}
}


2.模板2，可变滑动窗口：可变窗口的核心动作只有两个：向右扩大窗口，向左缩小窗口，双指针往右走，一般是right主动扩，left被动缩
func pattern2(nums []int, k int) {
 	left := 0

 for right := 0; right < len(nums); right++ {
      这里通常需要配合hash表，count := map[byte]int{}
     1. s[right] 进入窗口
     add(s[right])

     2. 如果窗口不合法，就不断缩小
     for windowIsInvalid() {

         remove(s[left])
         left++
     }

     3. 此时窗口合法
     ans = max(ans, right-left+1)
 }
 }

| 类型         | right 做什么 | left 什么时候动 | 什么时候更新答案 |
| ------------ | ------------ | --------------- | ---------------- |
| 固定窗口     | 不断加入     | 窗口超过 k      | 窗口大小等于 k   |
| 最长合法窗口 | 不断加入     | 窗口非法        | 窗口恢复合法后   |
| 最短满足窗口 | 不断加入     | 窗口满足条件    | 缩小过程中       |

1. 求最长合法窗口
for right := 0; right < n; right++ {

    加入 right

    for 窗口非法 {
        删除 left
        left++
    }

    ans = max(ans, right-left+1)
}
2. 求最短满足条件窗口
for right := 0; right < n; right++ {

    加入 right

    for 窗口已经满足条件 {

        更新最小答案

        删除 left
        left++
    }
}

