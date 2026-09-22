package main

// 算法思路：哈希表 + 双向链表。map 存 key→节点 用于 O(1) 定位；
// 双向链表按最近使用排序（head 端最新、tail 端最旧），Get/Put 命中后统一 moveToHead，
// 超容量时淘汰 tail.prev 并同步删除 map 中的键
type LRUCache struct {
    capacity int
    cache map[int]*Node
    head *Node
    tail *Node
    
}
type Node struct {
    key int
    val int
    prev *Node
    next *Node
}


func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head
    return LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		head: head,
		tail: tail,
    } 
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}
	node := &Node{
		key: key,
		val: value,
	}
	this.cache[key] = node
	this.addToHead(node)
	if len(this.cache) > this.capacity {
		removed := this.tail.prev
		this.remove(removed)
		delete(this.cache,removed.key)
	}    
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

}
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node

}
func (this *LRUCache) moveToHead(node *Node) {
	this.remove(node)
	this.addToHead(node)
}


func main(){

}