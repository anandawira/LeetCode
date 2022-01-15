package shopee

// type LRUCache struct {
// 	head, tail *Node
// 	mp         map[int]*Node
// 	capacity   int
// }

// type Node struct {
// 	prev, next *Node
// 	key, value int
// }

// func Constructor(capacity int) LRUCache {
// 	head := Node{
// 		prev:  nil,
// 		next:  nil,
// 		key:   0,
// 		value: 0,
// 	}

// 	tail := Node{
// 		prev:  nil,
// 		next:  nil,
// 		key:   0,
// 		value: 0,
// 	}

// 	head.next = &tail
// 	tail.prev = &head

// 	return LRUCache{
// 		head:     &head,
// 		tail:     &tail,
// 		mp:       make(map[int]*Node),
// 		capacity: capacity,
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	if n, ok := this.mp[key]; ok {
// 		this.remove(n)
// 		this.insert(n)
// 		return n.value
// 	}

// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	if _, ok := this.mp[key]; ok {
// 		this.remove(this.mp[key])
// 	}

// 	if len(this.mp) == this.capacity {
// 		this.remove(this.tail.prev)
// 	}

// 	this.insert(&Node{
// 		prev:  nil,
// 		next:  nil,
// 		key:   key,
// 		value: value,
// 	})
// }

// func (this *LRUCache) remove(node *Node) {
// 	delete(this.mp, node.key)
// 	node.prev.next = node.next
// 	node.next.prev = node.prev
// }

// func (this *LRUCache) insert(node *Node) {
// 	this.mp[node.key] = node
// 	next := this.head.next
// 	this.head.next = node
// 	node.prev = this.head
// 	next.prev = node
// 	node.next = next
// }
