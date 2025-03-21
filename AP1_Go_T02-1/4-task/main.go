package main

import (
	"fmt"
	"sync"
)

type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]*node[K, V]
	head     *node[K, V]
	tail     *node[K, V]
	mutex    sync.Mutex
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	if capacity <= 0 {
		panic("Capacity 0 yoki 0 dan kichkina bo'lishu mumkun emas !!!")

	}
	return &Cache[K, V]{
		capacity: capacity,
		items:    make(map[K]*node[K, V]),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if existingNode, found := c.items[key]; found {
		existingNode.value = value
		c.moveToFront(existingNode)
	}

	newNode := &node[K, V]{key: key, value: value}
	c.items[key] = newNode
	c.insertToFront(newNode)

	if len(c.items) > c.capacity {
		c.removeOldest()
	}

}

func (c *Cache[K, V]) removeOldest() {
	if c.tail == nil {
		return
	}
	delete(c.items, c.tail.key)
	c.detach(c.tail)
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if existingNode, found := c.items[key]; found {
		c.moveToFront(existingNode)
		return existingNode.value, true
	}
	var zeroValue V
	return zeroValue, false
}

func (c *Cache[K, V]) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items = make(map[K]*node[K, V])
	c.head = nil
	c.tail = nil
}

func (c *Cache[K, V]) moveToFront(n *node[K, V]) {
	if n == c.head {
		return
	}
	c.detach(n)
	c.insertToFront(n)
}

func (c *Cache[K, V]) detach(n *node[K, V]) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		c.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		c.tail = n.prev
	}
}

func (c *Cache[K, V]) insertToFront(n *node[K, V]) {
	if c.head == nil {
		c.head, c.tail = n, n
		return
	}

	n.next = c.head
	c.head.prev = n
	c.head = n
}

func main() {
	cache := NewCache[int, string](5)

	cache.Set(1, "Bir")
	cache.Set(2, "Ikki")
	cache.Set(2, "Uch")
	cache.Set(3, "Uch")

	fmt.Println(cache.Get(2))

}
