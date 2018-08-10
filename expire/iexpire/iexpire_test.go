package iexpire

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExipreMapCommon(t *testing.T) {
	key1 := "key1"
	key2 := "key2"
	key3 := "key3"

	e1 := "value1"
	e2 := "value2"
	e3 := "value3"

	d := time.Millisecond * 100

	em := NewExpireMap(time.Second)

	em.Set(key1, e1, d)
	em.SetExpiredIn(key2, e2, time.Now().Add(d))
	assert.Equal(t, e1, em.Get(key1))
	assert.True(t, em.Has(key1))
	assert.Equal(t, e2, em.Get(key2))
	assert.True(t, em.Has(key2))
	assert.Equal(t, 2, em.Size())

	em.Update(key1, e3)
	assert.Equal(t, e3, em.Get(key1), "update should work")

	time.Sleep(d)
	assert.False(t, em.Has(key1))
	assert.Nil(t, em.Get(key1), "should remove expired")
	assert.False(t, em.Has(key2))
	assert.Nil(t, em.Get(key2), "should remove expired")
	assert.Equal(t, 0, em.Size())

	// update un exists key
	assert.Nil(t, em.Get(key3))
	assert.Equal(t, NO_KEY_TO_UPDATE, em.Update(key3, e3))
}

func TestExipreMapGc(t *testing.T) {
	key1 := "key1"
	key2 := "key2"

	e1 := "value1"
	e2 := "value2"

	d := time.Millisecond * 200
	gcInterval := time.Millisecond * 500

	em := NewExpireMap(gcInterval)
	em.Set(key1, e1, d)
	em.Set(key2, e2, d)
	em.Gc()
	em.Gc()
	assert.Equal(t, 2, em.Size())
	time.Sleep(d)

	em.Gc()
	assert.Equal(t, 0, em.Size())

	// test auto gc
	em.Set(key1, e1, d)
	em.Set(key2, e2, d)
	time.Sleep(gcInterval)
	assert.Equal(t, 0, em.Size())

	// test clean up stop gc
	em.CleanUp()
	em.Set(key1, e1, d)
	em.Set(key2, e2, d)
	time.Sleep(gcInterval)
	assert.Equal(t, 2, em.Size())

	// test gc busy
	em2 := NewExpireMap(time.Millisecond * 200)
	em2.Set(key1, e1, d)
	em2.Set(key2, e2, d)
	em2.inGc = true
	time.Sleep(time.Millisecond * 200 * 2)
	assert.Equal(t, 2, em2.Size())
}
