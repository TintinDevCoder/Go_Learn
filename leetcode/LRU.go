package main

// 146. LRU 缓存
type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	cache    map[int]*Node
}
type Node struct {
	key, value int
	pre, next  *Node
}

func LRUCacheConstructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{capacity: capacity, head: head, tail: tail, cache: make(map[int]*Node)}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.cache[key]
	if ok {
		this.moveToHead(value)
		return value.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.cache[key]
	if ok {
		v.value = value
		this.moveToHead(v)
		return
	}
	if len(this.cache) >= this.capacity {
		delete(this.cache, this.tail.pre.key)
		removeNode(this.tail.pre)
	}
	node := &Node{key: key, value: value, pre: this.head, next: this.head.next}
	this.head.next.pre = node
	this.head.next = node
	this.cache[key] = node
}

func (this *LRUCache) moveToHead(node *Node) {
	removeNode(node)
	head := this.head
	node.next = head.next
	node.pre = head
	head.next.pre = node
	head.next = node
}

func removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
