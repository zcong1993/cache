package lru

import (
	"github.com/zcong1993/cache/lru/byteslru"
	"github.com/zcong1993/cache/lru/ilru"
	"github.com/zcong1993/cache/lru/stringlru"
)

//go:generate mkdir -p stringlru
//go:generate go_generics -i internal/lru/lru.go -t T=string -o stringlru/stringlru.go -p stringlru
//go:generate go_generics -i internal/lru/lru_test.go -t P=string -o stringlru/stringlru_test.go -p stringlru

// NewStringLru return a string lru instance
func NewStringLru(size int) *stringlru.Lru {
	return stringlru.NewLru(size)
}

//go:generate mkdir -p byteslru
//go:generate go_generics -i internal/lru/lru.go -t T=[]byte -o byteslru/byteslru.go -p byteslru
//go:generate go_generics -i internal/lru/lru_test.go -t P=[]byte -o byteslru/byteslru_test.go -p byteslru

// NewBytesLru return a bytes lru instance
func NewBytesLru(size int) *byteslru.Lru {
	return byteslru.NewLru(size)
}

// NewLru return a lru instance with type interface{}
func NewLru(size int) *ilru.Lru {
	return ilru.NewLru(size)
}
