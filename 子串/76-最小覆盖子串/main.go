package main

// 算法思路：滑动窗口 + 计数表。right 扩张把字符纳入窗口，count 记录已满足数量要求的字符种类数；
// 当 count == len(need) 说明窗口已覆盖 t，此时收缩 left 并记录最短区间，直到不再覆盖为止
func minWindow(s string, t string) string {
    need := make(map[byte]int)
    for i:=0;i<len(t);i++{
        need[t[i]]++
    }
    window := make(map[byte]int)
    left:=0
    count:=0
    start:=0
    length:=len(s)+1
    for right:=0;right<len(s);right++{
        c:=s[right]
        window[c]++
        if need[c]>0 && window[c]==need[c]{
            count++
        }
        for count==len(need){
            if right-left+1 < length{
                start=left
                length=right-left+1
            }
            d:=s[left]
            window[d]--
            if need[d]>0 && window[d]<need[d]{
                count--
            }
            left++
        }
    }
    if length==len(s)+1{
        return ""
    }
    return s[start:start+length]
}

func main(){

}