package main

  type Node struct {
      Val int
     Next *Node
      Random *Node
  }
 
// 算法思路：哈希表两遍遍历。第一遍为每个原节点建一个同值新节点，并记录 原节点→新节点 的映射；
// 第二遍借助映射把新节点的 Next / Random 指向各自对应的新节点
func copyRandomList(head *Node) *Node {
      m := make(map[*Node]*Node)
      curr := head 
      for curr != nil {
        m[curr] = &Node{curr.Val,nil,nil}
        curr = curr.Next
      }
      for old, new := range m {
          new.Next = m[old.Next]
          new.Random = m[old.Random]
      }
      return m[head]
    
}

func main(){

}