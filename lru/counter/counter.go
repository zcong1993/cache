package counter

import (
	"github.com/zcong1993/cache/lru"
	"github.com/zcong1993/cache/lru/int64lru"
)

// Counter is counter struct.
type Counter struct {
	cache *int64lru.Lru
}

// NewCounter is instance of counter.
func NewCounter(size int) *Counter {
	return &Counter{
		cache: lru.NewInt64Lru(size),
	}
}

// Incr incr a count by key, if not exists, set the count to 1.
func (c *Counter) Incr(key interface{}) int64 {
	v, ok := c.cache.Get(key)
	vv := int64(1)
	if ok {
		vv = *(v) + 1
	}
	c.cache.Add(key, vv)
	return vv
}

// Decr decr a count by key.
func (c *Counter) Decr(key interface{}) int64 {
	v, ok := c.cache.Get(key)
	if !ok {
		return 0
	}
	cc := *v - 1
	if cc <= 0 {
		c.cache.Remove(key)
		return 0
	}
	c.cache.Add(key, cc)
	return cc
}

// Get get counter by key, return 0 if key not exists.
func (c *Counter) Get(key interface{}) int64 {
	cc, ok := c.cache.Get(key)
	if !ok {
		return 0
	}
	return *cc
}

// Clear clear inner lru cache.
func (c *Counter) Clear() {
	c.cache.Clear()
}

// Len return inner lru size.
func (c *Counter) Len() int {
	return c.cache.Len()
}
