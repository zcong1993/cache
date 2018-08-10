package int64lru

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

func TestLru_Add_Get(t *testing.T) {
	var e1, e2, e3 int64
	e := createRandomObject(e1)
	if v, ok := e.(int64); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(int64); ok {
		e2 = v
	}
	e = createRandomObject(e3)
	if v, ok := e.(int64); ok {
		e3 = v
	}

	l := NewLru(1)
	l.Add("a", e1)
	v, ok := l.Get("a")
	assert.Equal(t, e1, *v)
	assert.True(t, ok)

	l.Add("b", e2)
	v, ok = l.Get("b")
	assert.True(t, ok)
	assert.Equal(t, e2, *v)

	_, ok = l.Get("a")
	assert.False(t, ok)

	l.Add("b", e3)
	v, ok = l.Get("b")
	assert.True(t, ok)
	assert.Equal(t, e3, *v)
}

func TestLru_Len(t *testing.T) {
	var e1, e2, e3 int64
	e := createRandomObject(e1)
	if v, ok := e.(int64); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(int64); ok {
		e2 = v
	}

	e = createRandomObject(e3)
	if v, ok := e.(int64); ok {
		e3 = v
	}

	l := NewLru(2)

	l.Add("a", e1)
	assert.Equal(t, 1, l.Len())

	l.Add("b", e2)
	assert.Equal(t, 2, l.Len())

	l.Add("c", e3)
	assert.Equal(t, 2, l.Len())
}

func TestLru_Remove(t *testing.T) {
	var e1 int64
	e := createRandomObject(e1)
	if v, ok := e.(int64); ok {
		e1 = v
	}

	l := NewLru(1)
	l.Add("a", e1)
	v, ok := l.Get("a")
	assert.Equal(t, e1, *v)
	assert.True(t, ok)

	l.Remove("a")
	_, ok = l.Get("a")
	assert.False(t, ok)
}

func TestLru_RemoveOldest(t *testing.T) {
	var e1, e2, e3 int64
	e := createRandomObject(e1)
	if v, ok := e.(int64); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(int64); ok {
		e2 = v
	}

	e = createRandomObject(e3)
	if v, ok := e.(int64); ok {
		e3 = v
	}

	l := NewLru(3)
	l.Add("a", e1)
	l.Add("b", e2)
	l.Add("c", e3)

	l.RemoveOldest()
	_, ok := l.Get("a")
	assert.False(t, ok)

	l.Get("b")
	l.RemoveOldest()

	_, ok = l.Get("c")
	assert.False(t, ok)
}

func TestLru_Clear(t *testing.T) {
	var e1, e2, e3 int64
	e := createRandomObject(e1)
	if v, ok := e.(int64); ok {
		e1 = v
	}
	e = createRandomObject(e2)
	if v, ok := e.(int64); ok {
		e2 = v
	}

	e = createRandomObject(e3)
	if v, ok := e.(int64); ok {
		e3 = v
	}

	l := NewLru(3)
	l.Add("a", e1)
	l.Add("b", e2)
	l.Add("c", e3)

	assert.Equal(t, 3, l.Len())

	l.Clear()

	assert.Equal(t, 0, l.Len())
}

func createRandomObject(i interface{}) interface{} {
	v, ok := quick.Value(reflect.TypeOf(i), rand.New(rand.NewSource(time.Now().UnixNano())))
	if !ok {
		panic(fmt.Sprintf("unsupported type %v", i))
	}
	return v.Interface()
}
