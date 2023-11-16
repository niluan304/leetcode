package main

import "container/list"

type LRUCache struct {
	capacity int                   // 容量
	list     *list.List            // 双向链表
	data     map[int]*list.Element // 缓存数据
}

type Value struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		data:     make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) int {
	e, ok := c.data[key]
	// 不存在
	if !ok {
		return -1
	}

	// 更新使用的key
	c.list.MoveToFront(e)
	return e.Value.(*Value).value
}

func (c *LRUCache) Put(key int, value int) {
	v := &Value{
		key:   key,
		value: value,
	}

	if e, ok := c.data[key]; ok {
		// 更新使用的key
		c.list.MoveToFront(e)
		e.Value = v
	} else {
		e := c.list.PushFront(v)
		c.data[key] = e
		c.RemoveOld()
	}
}

// RemoveOld 逐出/淘汰 最久未使用的关键字
func (c *LRUCache) RemoveOld() {
	for len(c.data) > c.capacity {
		back := c.list.Back()
		if back == nil {
			return
		}

		c.list.Remove(back)
		old := (back.Value).(*Value)
		delete(c.data, old.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
