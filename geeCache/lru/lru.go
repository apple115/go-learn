package geeCache

import "container/list"

type Cache struct {
	maxBytest int64
	nbytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicated func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytest int64, onEvicated func(key string, value Value)) *Cache {
	return &Cache{
		maxBytest: maxBytest,
		nbytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicated: onEvicated,
	}
}

// 查找主要有 2 个步骤，第一步是从字典中找到对应的双向链表的节点，第二步，将该节点移动到队首
// Get ...
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 这里的删除，实际上是缓存淘汰。即移除最近最少访问的节点（队首）
// RemoveOldest ...
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= (int64(len(kv.key)) + int64(kv.value.Len()))
		if c.OnEvicated != nil {
			c.OnEvicated(kv.key, kv.value)
		}
	}
}

// Add ...
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytest != 0 && c.maxBytest < c.nbytes {
		c.RemoveOldest()
	}
}

// Len ...
func (c *Cache) Len() int {
	return c.ll.Len()
}
