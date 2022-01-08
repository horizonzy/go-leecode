package main

func main() {
	cache := Constructor(2)
	cache.Get(2)

	cache.Put(2, 6)
	cache.Get(1)

	cache.Put(1, 5)
	cache.Put(1, 2)
	cache.Get(1)
	cache.Get(2)

}

type LRUCache struct {
	data     map[int]*Node
	head     *Node
	tail     *Node
	capacity int
	size     int
}

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{capacity: capacity, data: map[int]*Node{}, head: &Node{}, tail: &Node{}}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (cache *LRUCache) Get(key int) int {
	node := cache.data[key]
	if node == nil {
		return -1
	}
	if cache.size == 1 {
		return node.val
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	cache.head.next.prev = node
	node.next = cache.head.next

	cache.head.next = node
	node.prev = cache.head
	return node.val
}

func (cache *LRUCache) Put(key int, value int) {
	_, ok := cache.data[key]
	if !ok {
		cache.size++
		if cache.size > cache.capacity {
			delete(cache.data, cache.tail.prev.key)
			cache.size--
			if cache.size == 1 {
				cache.head.next = cache.tail
				cache.tail.prev = cache.head
			} else {
				cache.tail.prev.prev.next = cache.tail
				cache.tail.prev = cache.tail.prev.prev
			}
		}
		node := &Node{val: value, key: key}
		cache.data[key] = node

		cache.head.next.prev = node
		node.next = cache.head.next
		cache.head.next = node
		node.prev = cache.head
	} else {
		node := cache.data[key]
		node.val = value

		if cache.size == 1 {
			return
		}

		node.prev.next = node.next
		node.next.prev = node.prev

		cache.head.next.prev = node
		node.next = cache.head.next
		cache.head.next = node
		node.prev = cache.head
	}
}
