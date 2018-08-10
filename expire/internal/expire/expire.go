package expire

import (
	"errors"
	"sync"
	"time"
)

var (
	// NO_KEY_TO_UPDATE is error message
	NO_KEY_TO_UPDATE = errors.New("key not exists or expired, set it first. ")
)

// T is a Template type.
type T = Template

// ExipreMap is struct of expire map
type ExipreMap struct {
	store      map[interface{}]Value
	GcInterval time.Duration
	inGc       bool
	mu         *sync.RWMutex
	t          *time.Ticker
}

// Value is expire map value
type Value struct {
	Val       T
	expiredIn time.Time
}

// NewExpireMap is constructor of ExipreMap
func NewExpireMap(interval time.Duration) *ExipreMap {
	em := &ExipreMap{
		store:      make(map[interface{}]Value),
		GcInterval: interval,
		inGc:       false,
		mu:         new(sync.RWMutex),
	}
	go em.startGc()
	return em
}

func (em *ExipreMap) startGc() {
	t := time.NewTicker(em.GcInterval)
	em.t = t
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if em.inGc {
				return
			}
			em.Gc()
		}
	}
}

func isExpiredValue(val Value) bool {
	if val.expiredIn.UnixNano() <= time.Now().UnixNano() {
		return true
	}
	return false
}

// Gc is gc function
func (em *ExipreMap) Gc() {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.inGc = true
	for k, v := range em.store {
		if isExpiredValue(v) {
			delete(em.store, k)
		}
	}
	em.inGc = false
}

// Get implement map get method
func (em *ExipreMap) Get(key interface{}) *T {
	em.mu.RLock()
	defer em.mu.RUnlock()
	v, ok := em.store[key]
	if !ok {
		return nil
	}
	if isExpiredValue(v) {
		delete(em.store, key)
		return nil
	}
	return &v.Val
}

// Set implement map set function but with expire
func (em *ExipreMap) Set(k interface{}, val T, expire time.Duration) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.store[k] = Value{
		Val:       val,
		expiredIn: time.Now().Add(expire),
	}
}

// Set implement map set function but with expiredIn
func (em *ExipreMap) SetExpiredIn(k interface{}, val T, expiredIn time.Time) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.store[k] = Value{
		Val:       val,
		expiredIn: expiredIn,
	}
}

// Size show the map size, not strict
func (em *ExipreMap) Size() int {
	em.mu.RLock()
	defer em.mu.RUnlock()
	return len(em.store)
}

// Has implement map has function
func (em *ExipreMap) Has(k interface{}) bool {
	em.mu.RLock()
	defer em.mu.RUnlock()
	v, ok := em.store[k]
	if !ok {
		return false
	}
	return !isExpiredValue(v)
}

// Update update a item but not update expire time
func (em *ExipreMap) Update(k interface{}, newVal T) error {
	if !em.Has(k) {
		return NO_KEY_TO_UPDATE
	}
	em.mu.Lock()
	defer em.mu.Unlock()
	val := em.store[k]
	val.Val = newVal
	em.store[k] = val
	return nil
}

// CleanUp stop auto gc
func (em *ExipreMap) CleanUp() {
	if em.t != nil {
		em.t.Stop()
	}
}
