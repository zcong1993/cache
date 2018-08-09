package ilru

import (
	"container/list"
)

type entry struct {
	key   interface{}
	value interface{}
}

// Lru is lru cache struct.
type Lru struct {
	ll    *list.List
	store map[interface{}]*list.Element
	size  int
}

// NewLru create and init lru instance.
func NewLru(size int) *Lru {
	return &Lru{
		ll:    list.New(),
		store: make(map[interface{}]*list.Element),
		size:  size,
	}
}

// Add add a key value to cache.
func (l *Lru) Add(key interface{}, v interface{}) {
	// update if already exists
	if ent, exists := l.store[key]; exists {
		ent.Value.(*entry).value = v
		l.ll.MoveToFront(ent)
		return
	}
	// add new
	ele := l.ll.PushFront(&entry{key: key, value: v})
	l.store[key] = ele
	// check is full
	if l.ll.Len() > l.size {
		// remove oldest
		l.RemoveOldest()
	}
}

// Remove remove a key from cache.
func (l *Lru) removeElement(ele *list.Element) {
	l.ll.Remove(ele)
	k := ele.Value.(*entry).key
	delete(l.store, k)
}

// Get get cache by key.
func (l *Lru) RemoveOldest() {
	ele := l.ll.Back()
	if ele != nil {
		l.removeElement(ele)
	}
}

// Len return cache size.
func (l *Lru) Remove(k interface{}) {
	if ele, exists := l.store[k]; exists {
		l.removeElement(ele)
	}
}

// Clear flush the cache.
func (l *Lru) Get(k interface{}) (interface{}, bool) {
	if ele, exists := l.store[k]; exists {
		l.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return nil, false
}

func (l *Lru) Len() int {
	return l.ll.Len()
}

func (l *Lru) Clear() {
	ll := NewLru(l.size)
	l.store = ll.store
	l.ll = ll.ll
}
