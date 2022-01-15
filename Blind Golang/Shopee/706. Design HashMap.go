package shopee

type Node struct {
	key, val int
	next     *Node
}

type MyHashMap struct {
	items []*Node
	size  int
}

func Constructor() MyHashMap {
	const size = 1000
	return MyHashMap{
		items: make([]*Node, size),
		size:  size,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	index := key % this.size

	if this.items[index] == nil {
		this.items[index] = &Node{key: key, val: value}
		return
	}

	item := this.items[index]

	for item != nil {
		if item.key == key {
			item.val = value
			return
		}

		if item.next == nil {
			item.next = &Node{key: key, val: value}
			return
		}

		item = item.next
	}
}

func (this *MyHashMap) Get(key int) int {
	index := key % this.size
	if this.items[index] == nil {
		return -1
	}

	item := this.items[index]

	for item != nil {
		if item.key == key {
			return item.val
		}

		item = item.next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := key % this.size
	if this.items[index] == nil {
		return
	}

	item := this.items[index]

	if item.key == key {
		this.items[index] = item.next
	}

	prev, item := item, item.next

	for item != nil {
		if item.key == key {
			prev.next = item.next
			return
		}

		prev, item = item, item.next
	}
}
