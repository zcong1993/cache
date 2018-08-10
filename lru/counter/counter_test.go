package counter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounterIncr(t *testing.T) {
	key1 := "a"
	key2 := "b"

	c := NewCounter(1)
	assert.Equal(t, int64(1), c.Incr(key1))
	assert.Equal(t, int64(2), c.Incr(key1))

	assert.Equal(t, int64(1), c.Incr(key2))
	assert.Equal(t, int64(1), c.Incr(key1), "key1 should be reset cause counter lru size is only 1")
}

func TestCounterDecr(t *testing.T) {
	key1 := "a"

	c := NewCounter(1)
	c.Incr(key1)
	c.Incr(key1)

	c.Decr(key1)
	assert.Equal(t, int64(1), c.Get(key1))
	c.Decr(key1)
	assert.Equal(t, int64(0), c.Get(key1))
	c.Decr(key1)
	assert.Equal(t, int64(0), c.Get(key1))
}

func TestCounterLen(t *testing.T) {
	c := NewCounter(2)

	c.Incr("a")
	assert.Equal(t, 1, c.Len())
	c.Incr("b")
	assert.Equal(t, 2, c.Len())
	c.Incr("c")
	assert.Equal(t, 2, c.Len())
}

func TestCounterClear(t *testing.T) {
	c := NewCounter(2)

	c.Incr("a")
	c.Incr("b")

	c.Clear()
	assert.Equal(t, 0, c.Len())
}
