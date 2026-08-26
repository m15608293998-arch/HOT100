package main

import "fmt"

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	newNode := &Node{
		key:   key,
		value: value,
	}
	this.cache[key] = newNode
	this.addToHead(newNode)

	if len(this.cache) > this.capacity {
		removed := this.removeTail()
		delete(this.cache, removed.key)
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// 演示用例（LeetCode 146 官方示例）
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 1
	cache.Put(3, 3)           // 淘汰 key 2
	fmt.Println(cache.Get(2)) // -1
	cache.Put(4, 4)           // 淘汰 key 1
	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}
