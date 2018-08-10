package lru

import (
	"github.com/zcong1993/cache/lru/byteslru"
	"github.com/zcong1993/cache/lru/ilru"
	"github.com/zcong1993/cache/lru/int64lru"
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

//go:generate mkdir -p int64lru
//go:generate go_generics -i internal/lru/lru.go -t T=int64 -o int64lru/int64lru.go -p int64lru
//go:generate go_generics -i internal/lru/lru_test.go -t P=int64 -o int64lru/int64lru_test.go -p int64lru

// NewInt64Lru an int64 lru instance
func NewInt64Lru(size int) *int64lru.Lru {
	return int64lru.NewLru(size)
}

// NewLru return a lru instance with type interface{}
func NewLru(size int) *ilru.Lru {
	return ilru.NewLru(size)
}
